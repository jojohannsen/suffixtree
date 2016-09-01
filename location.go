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
	edge          *Edge
	OnNode        bool
	OffsetFromTop int64
	Base          Node
}

func NewLocation(node Node) *Location {
	return &Location{nil, true, 0, node}
}

func (location *Location) String() string {
	if location == nil {
		return "Location=NIL"
	}
	edge := "nil"
	if location.edge != nil {
		edge = fmt.Sprintf("%s", location.edge)
	}
	node := "nil"
	if location.Base != nil {
		node = fmt.Sprintf("%s", location.Base)
	}
	return fmt.Sprintf("Location(%s, %t, %d, %s)",
		edge, location.OnNode, location.OffsetFromTop, node)
}
