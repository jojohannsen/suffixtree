package suffixtree

type Builder interface {
	split(parent, child Node, edge *Edge, splitOffset int32) Node
}

type builder struct {
	_idFactory *idFactory
	dataSource DataSource
}

func NewBuilder(_idFactory *idFactory, dataSource DataSource) Builder {
	return &builder{_idFactory, dataSource}
}

// split an Edge, return the newly created Node
func (b *builder) split(parent, child Node, edge *Edge, splitOffset int32) Node {
	topEdge := edge
	bottomEdge := NewEdge(topEdge.StartOffset+splitOffset, topEdge.EndOffset)
	topEdge.EndOffset = bottomEdge.StartOffset - 1
	internalNode := NewInternalNode(b._idFactory.NextId(), parent, topEdge)
	parent.AddOutgoingEdgeNode(b.dataSource.KeyAtOffset(topEdge.StartOffset), topEdge, internalNode)
	internalNode.AddOutgoingEdgeNode(b.dataSource.KeyAtOffset(bottomEdge.StartOffset), bottomEdge, child)
	child.setIncoming(internalNode, bottomEdge)
	return internalNode
}
