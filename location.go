package suffixtree

import "fmt"

// Track location in suffix tree while building
// Either on a node:
//   onNode: true
//   node: <node we are on>
//   edge: N/A
//   offsetFromTop:  N/A
//
// Or on an edge:
//   onNode: false
//   node: <node below edge we are on>
//   edge: <edge we are on>
//   offsetFromTop: how many characters into edge (between 1 and edge length-1)
//
type Location struct {
	Edge          *Edge
	OnNode        bool
	OffsetFromTop int32
	Base          Node
}

func NewLocation(node Node) *Location {
	return &Location{nil, true, 0, node}
}

func (location *Location) String() string {
	if location == nil {
		return "Location=NIL"
	}

	if location.OnNode {
		return fmt.Sprintf("Location on Node %d", location.Base.Id())
	} else {
		return fmt.Sprintf("Location on Edge, OffsetFromTop=%d, node is %d",
			location.OffsetFromTop, location.Base.Id())
	}
}
