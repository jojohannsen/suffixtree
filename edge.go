package suffixtree

import "fmt"

// An Edge connects two nodes.  Trees are directional -- the Node above the Edge refers to it as a outgoingEdge,
// the Node below the Edge refers to it as a incomingEdge

type Edge struct {
	StartOffset int64
	EndOffset   int64
}

const FinalOffset int64 = -1
const EdgeTerminatesAtEnd int64 = -2

func NewEdge(startOffset, endOffset int64) *Edge {
	return &Edge{startOffset, endOffset}
}

func NewLeafEdge(startOffset int64) *Edge {
	return &Edge{startOffset, FinalOffset}
}

func (edge *Edge) String() string {
	return fmt.Sprintf("[%d,%d]", edge.StartOffset, edge.EndOffset)
}

func (edge *Edge) length() int64 {
	if edge.EndOffset == FinalOffset {
		return EdgeTerminatesAtEnd
	}
	return edge.EndOffset - edge.StartOffset + 1
}
