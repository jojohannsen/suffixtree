package suffixtree

import "fmt"

type Ukkonen interface {
	Extend() bool
	Finish()
	Debug(dChan chan string)
	DrainDataSource()
	Tree() SuffixTree
	Location() *Location
	DataSource() DataSource
}

type ukkonen struct {
	dataChannel     <-chan STKey
	offset          int64
	location        *Location
	root            Node
	suffixTree      SuffixTree
	dataSource      DataSource
	needsSuffixLink Node
	builder         Builder
	traverser       Traverser
	debugChannel    chan string
}

func (b *ukkonen) Debug(dChan chan string) {
	b.debugChannel = dChan
}

func (b *ukkonen) DataSource() DataSource {
	return b.dataSource
}

func (b *ukkonen) Location() *Location {
	return b.location
}

func NewUkkonen(dataSource DataSource) Ukkonen {
	suffixTree := NewSuffixTree(dataSource)
	root := suffixTree.Root()
	return &ukkonen{dataSource.STKeys(), 0, NewLocation(root), root,
		suffixTree, dataSource, nil,
		NewBuilder(dataSource), NewTraverser(dataSource), nil}
}

func (b *ukkonen) DrainDataSource() {
	for b.Extend() {
	}
}

// extend the suffix tree with the next value, returns false if data channel is closed
func (b *ukkonen) Extend() bool {
	value, ok := <-b.dataChannel
	if !ok {
		return false
	}

	// increment the offset after each successful read
	defer func(b *ukkonen) {
		b.offset++
		//treePrintWithTitle(fmt.Sprintf("FINISHED extend(%d)", value), b.root, b.location)
	}(b)

	// otherwise, extend until done
	if b.debugChannel != nil {
		b.debugChannel <- fmt.Sprintf("Extend with value '%s' at %s", string(value), b.location)
	}

	b.finish(value)
	if b.debugChannel != nil {
		b.debugChannel <- fmt.Sprintf("Done with extension for '%s'", string(value))
	}
	return true
}

func (b *ukkonen) prepareForNextExtension() {
	if b.debugChannel != nil {
		b.debugChannel <- fmt.Sprintf(" prepareForNextExtension starting at %s", b.location)
	}
	b.traverser.traverseToNextSuffix(b.location, b.debugChannel)
	if b.debugChannel != nil {
		b.debugChannel <- fmt.Sprintf(" .. after traverseToNextSuffix %s", b.location)
	}
	// if we are on the root, and there's a node needing a suffix link, set it
	// if it's not the root, we will be creating it here, and have to set it
	// after it gets created
	if b.location.Base.isRoot() && b.needsSuffixLink != nil {
		b.needsSuffixLink.SetSuffixLink(b.location.Base)
		b.needsSuffixLink = nil
	}
}

func (b *ukkonen) finish(value STKey) {
	for b.extendWithValue(value) {
		b.prepareForNextExtension()
	}
}

func (b *ukkonen) Finish() {
	b.finish(STKey('$'))
}

func (b *ukkonen) extendWithValue(value STKey) bool {
	if b.location.OnNode {
		if b.debugChannel != nil {
			b.debugChannel <- fmt.Sprintf(" extendWithValue on Node %d", b.location.Base.Id())
		}
		// if the previous node needs a suffix link, this is the place
		if b.needsSuffixLink != nil {
			b.needsSuffixLink.SetSuffixLink(b.location.Base)
			b.needsSuffixLink = nil
		}
		// if child value is there, just update location
		if b.location.Base.EdgeFollowing(value) != nil {
			if b.debugChannel != nil {
				b.debugChannel <- fmt.Sprintf("   node has an edge with that value, nothing to do")
			}
			b.traverser.traverseOne(b.location, value)
			return false
		} else {
			// otherwise we add the value
			edge, node := b.location.Base.addLeafEdgeNode(value, b.offset)
			if b.debugChannel != nil {
				b.debugChannel <- fmt.Sprintf("   creating leaf edge, new node is %d, edge %s", node.Id(), edge)
			}
			b.location.OnNode = true
			return !b.location.Base.isRoot()
		}

	} else {
		if b.debugChannel != nil {
			b.debugChannel <- fmt.Sprintf(" extendWithValue on Edge, offset %d ending at Node %d", b.location.OffsetFromTop, b.location.Base.Id())
		}
		// we are on the edge, see if the character at the offset matches
		valueOffset := b.location.Edge.StartOffset + b.location.OffsetFromTop
		if b.dataSource.keyAtOffset(valueOffset) == value {
			// Rule 3, value already in tree, change location and we are done
			b.location.OffsetFromTop++
			if b.location.OffsetFromTop == b.location.Edge.length() {
				b.location.OnNode = true
			}
			return false
		} else if b.location.Base.isRoot() {
			// add leaf, set location
			leafEdge, leafNode := NewLeafEdgeNode(b.root, b.offset)
			b.location.Base.AddOutgoingEdgeNode(value, leafEdge, leafNode)
			b.location.Base = leafNode
			b.location.OffsetFromTop = 0
		} else {
			// Rule 2
			// - split the edge, and let builder know the new Node needs a suffix link on the next extension
			previousNeedsSuffixLink := b.needsSuffixLink
			b.needsSuffixLink = b.builder.split(b.location.Base.parent(), b.location.Base, b.location.Edge, b.location.OffsetFromTop)
			if previousNeedsSuffixLink != nil {
				previousNeedsSuffixLink.SetSuffixLink(b.needsSuffixLink)
			}
			if b.debugChannel != nil {
				b.debugChannel <- fmt.Sprintf(" extendWithValue split edge, creating Node %d", b.needsSuffixLink.Id())
			}
			// - add the new leaf node
			leafEdge, leafNode := NewLeafEdgeNode(b.needsSuffixLink, b.offset)
			b.needsSuffixLink.AddOutgoingEdgeNode(value, leafEdge, leafNode)

			// after the split, we are located on the internal node
			b.location.Base = b.needsSuffixLink
			b.location.OnNode = true

			// - return true so we continue extending
			return true
		}
	}
	return false
}

func (b *ukkonen) Tree() SuffixTree {
	return b.suffixTree
}
