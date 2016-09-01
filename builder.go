package suffixtree

type Builder interface {
	split(parent, child Node, edge *Edge, splitOffset int64) Node
}

type builder struct {
	dataSource DataSource
}

func NewBuilder(dataSource DataSource) Builder {
	return &builder{dataSource}
}

// split an Edge, return the newly created Node
func (b *builder) split(parent, child Node, edge *Edge, splitOffset int64) Node {
	topEdge := edge
	bottomEdge := NewEdge(topEdge.StartOffset+splitOffset, topEdge.EndOffset)
	topEdge.EndOffset = bottomEdge.StartOffset - 1
	internalNode := NewInternalNode(parent, topEdge)
	parent.addOutgoingEdgeNode(b.dataSource.keyAtOffset(topEdge.StartOffset), topEdge, internalNode)
	internalNode.addOutgoingEdgeNode(b.dataSource.keyAtOffset(bottomEdge.StartOffset), bottomEdge, child)
	child.setIncoming(internalNode, bottomEdge)
	return internalNode
}
