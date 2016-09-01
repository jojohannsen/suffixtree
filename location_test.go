package suffixtree

import (
	"fmt"
	"testing"
)

func TestLocationString(t *testing.T) {
	location := NewLocation(NewRootNode())
	want := "Location(nil, true, 0, ROOT())"
	s := fmt.Sprintf("%s", location)
	if s != want {
		t.Errorf("NewLocation() got %s, want %s", s, want)
	}
	location.edge = NewEdge(10, 20)
	s = fmt.Sprintf("%s", location)
	want = "Location([10,20], true, 0, ROOT())"
	if s != want {
		t.Errorf("After setting edge, got %s, want %s", s, want)
	}
	location.OffsetFromTop = 9999
	want = "Location([10,20], true, 9999, ROOT())"
	s = fmt.Sprintf("%s", location)
	if s != want {
		t.Errorf("After setting offsetFromTop, got %s, want %s", s, want)
	}
	location.Base = NewRootNode()
	s = fmt.Sprintf("%s", location)
	want = "Location([10,20], true, 9999, ROOT())"
	if s != want {
		t.Errorf("After setting node, got %s, want %s", s, want)
	}
	location.OnNode = false
	s = fmt.Sprintf("%s", location)
	want = "Location([10,20], false, 9999, ROOT())"
	if s != want {
		t.Errorf("After setting onInternalNode, got %s, want %s", s, want)
	}
}
