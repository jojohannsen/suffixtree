package suffixtree

import "fmt"

type Traverser interface {
	traverseToNextSuffix(location *Location)
	traverseOne(location *Location, value STKey)
	traverseDownValue(location *Location, value STKey) bool
}

type traverser struct {
	dataSource            DataSource
	numberValuesTraversed int64
	traversedDataOffset   int64
}

func NewTraverser(dataSource DataSource) Traverser {
	return &traverser{dataSource, 0, 0}
}

func (t *traverser) String() string {
	return fmt.Sprintf("(%d values starting at %d)", t.numberValuesTraversed, t.traversedDataOffset)
}

// traverse down a value, return true and location updated if value is present
func (t *traverser) traverseDownValue(location *Location, value STKey) bool {
	if location.OnNode {
		if location.Base.NodeFollowing(value) != nil {
			t.traverseOne(location, value)
			return true
		} else {
			return false
		}
	} else {
		offsetToCheck := location.Base.IncomingEdge().StartOffset + location.OffsetFromTop
		if t.dataSource.keyAtOffset(offsetToCheck) == value {
			t.traverseEdgeValue(location)
			return true
		}
	}
	return false
}

// set the Location to be at the next suffix
func (t *traverser) traverseToNextSuffix(location *Location) {
	t.traverseUp(location)
	t.traverseSuffixLink(location)
	t.traverseDown(location)
}

func (t *traverser) traverseUp(location *Location) {
	location.OnNode = true
	t.numberValuesTraversed = location.OffsetFromTop
	location.OffsetFromTop = 0
	t.traversedDataOffset = location.edge.StartOffset + location.OffsetFromTop
	location.Base = location.Base.parent()
	location.edge = location.Base.IncomingEdge()
}

func (t *traverser) traverseSuffixLink(location *Location) {
	if location.Base.isRoot() {
		t.numberValuesTraversed -= 1
		t.traversedDataOffset += 1
	} else {
		location.Base = location.Base.suffixLink()
	}
}

func (t *traverser) traverseOne(location *Location, value STKey) {
	location.edge = location.Base.edgeFollowing(value)
	location.Base = location.Base.NodeFollowing(value)
	location.OnNode = (location.edge.EndOffset != FinalOffset) && (location.edge.length() == 1)
	if (location.edge.EndOffset == FinalOffset) || (location.edge.length() > 1) {
		location.OffsetFromTop = 1
	}
}

func (t *traverser) traverseEdgeValue(location *Location) {
	location.OffsetFromTop++
	if location.OffsetFromTop == location.edge.length() {
		location.OnNode = true
		location.OffsetFromTop = 0
	}
}

// skip count down traversal
func (t *traverser) traverseDown(location *Location) {
	for t.numberValuesTraversed > 0 {
		location.edge, location.Base = location.Base.outgoingEdgeNode(t.dataSource.keyAtOffset(t.traversedDataOffset))
		edgeLength := location.edge.length()
		if (t.numberValuesTraversed < edgeLength) || (edgeLength == EdgeTerminatesAtEnd) {
			location.OffsetFromTop = t.numberValuesTraversed
			t.numberValuesTraversed = 0
			location.OnNode = false
		} else {
			location.OnNode = t.numberValuesTraversed == edgeLength
			t.numberValuesTraversed -= edgeLength
			t.traversedDataOffset += edgeLength
		}
	}
}
