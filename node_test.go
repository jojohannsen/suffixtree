package suffixtree

import (
	"strconv"
	"testing"
)

func TestRootNode(t *testing.T) {
	var root Node = NewRootNode()
	if root == nil {
		t.Error("NewRootNode returned nil")
	}
	if root.IncomingEdge() != nil {
		t.Error("root has unexpected parent")
	}
	if root.edgeFollowing(0) != nil {
		t.Error("unexpected children")
	}
	if root.suffixLink() != nil {
		t.Error("Unexpected suffix link")
	}
}

func TestRootOutgoingEdge(t *testing.T) {
	root := NewRootNode()
	edge := NewEdge(0, 3)
	root.addOutgoingEdgeNode(1000, edge, nil)
	testEdge := root.edgeFollowing(1000)
	if testEdge != edge {
		t.Error("Did not get edge just added")
	}
	testEdge = root.edgeFollowing(1001)
	if testEdge != nil {
		t.Error("Found edge where none expected")
	}
}

func TestRootPanic1(t *testing.T) {
	defer func() {
		recover()
	}()
	root := NewRootNode()
	root.setSuffixLink(nil)
	t.Errorf("The code did not panic")
}

func TestRootChildren(t *testing.T) {
	var rootTests = []struct {
		startOffset int64
		endOffset   int64
		firstValue  STKey
	}{
		{0, 3, 1001},
		{1, 3, 1002},
		{2, 3, 1003},
		{3, 3, 1004},
	}
	root := NewRootNode()
	for _, test := range rootTests {
		edge := NewEdge(test.startOffset, test.endOffset)
		root.addOutgoingEdgeNode(test.firstValue, edge, nil)
	}
	for _, test := range rootTests {
		testEdge := root.edgeFollowing(test.firstValue)
		if testEdge.StartOffset != test.startOffset {
			t.Error("Edge start offset, got " + strconv.FormatInt(testEdge.StartOffset, 10) + ", want " + strconv.FormatInt(test.startOffset, 10))
		}
		if testEdge.EndOffset != test.endOffset {
			t.Error("Edge end offset, got " + strconv.FormatInt(testEdge.EndOffset, 10) + ", want " + strconv.FormatInt(test.endOffset, 10))
		}
	}
	testEdge := root.edgeFollowing(111111)
	if testEdge != nil {
		t.Error("Found edge where none expected")
	}
}

func TestInternalChildren(t *testing.T) {
	var internalTests = []struct {
		startOffset int64
		endOffset   int64
		firstValue  STKey
	}{
		{0, 3, 1001},
		{1, 3, 1002},
		{2, 3, 1003},
		{3, 3, 1004},
	}
	root := NewRootNode()
	for _, test := range internalTests {
		edge := NewEdge(test.startOffset, test.endOffset)
		internal := NewInternalNode(root, edge)
		if internal.IncomingEdge() != edge {
			t.Error("Internal parent not found")
		}
		internalEdge := NewEdge(test.startOffset, test.endOffset)
		root.addOutgoingEdgeNode(test.firstValue, edge, nil)
		internal.addOutgoingEdgeNode(2000, internalEdge, nil)
		if internal.suffixLink() != nil {
			t.Error("new Internal Node had a suffix link")
		}
		if internal.edgeFollowing(2000) != internalEdge {
			t.Error("internal edge follow failed")
		}
		internal.setSuffixLink(root)
		if internal.suffixLink() != root {
			t.Error("InternalNode setSuffixLink failed")
		}
	}
	for _, test := range internalTests {
		testEdge := root.edgeFollowing(test.firstValue)
		if testEdge.StartOffset != test.startOffset {
			t.Error("Edge start offset, got " + strconv.FormatInt(testEdge.StartOffset, 10) + ", want " + strconv.FormatInt(test.startOffset, 10))
		}
		if testEdge.EndOffset != test.endOffset {
			t.Error("Edge end offset, got " + strconv.FormatInt(testEdge.EndOffset, 10) + ", want " + strconv.FormatInt(test.endOffset, 10))
		}
	}
	testEdge := root.edgeFollowing(111111)
	if testEdge != nil {
		t.Error("Found edge where none expected")
	}
}

func TestLeafOutgoingFollowing(t *testing.T) {
	defer func() {
		recover()
	}()

	root := NewRootNode()
	edge, leaf := NewLeafEdgeNode(root, 0)
	leaf.addOutgoingEdgeNode(1000, edge, leaf)
	t.Errorf("The code did not panic")
}

func TestLeafEdgeFollowingPanic(t *testing.T) {
	defer func() {
		recover()
	}()

	root := NewRootNode()
	_, leaf := NewLeafEdgeNode(root, 0)
	leaf.edgeFollowing(1000)
	t.Errorf("The code did not panic")
}

func TestLeafSuffixLinkPanic(t *testing.T) {
	defer func() {
		recover()
	}()

	root := NewRootNode()
	_, leaf := NewLeafEdgeNode(root, 0)
	leaf.setSuffixLink(nil)
	t.Errorf("The code did not panic")
}

func TestRootLeaf(t *testing.T) {
	var rootLeafTests = []struct {
		startOffset int64
		endOffset   int64
		firstValue  STKey
	}{
		{0, -1, 1001},
		{1, -1, 1002},
		{2, -1, 1003},
		{3, -1, 1004},
	}
	root := NewRootNode()
	for _, test := range rootLeafTests {
		edge := NewEdge(test.startOffset, test.endOffset)
		edge, leaf := NewLeafEdgeNode(root, test.startOffset)
		if leaf.IncomingEdge() != edge {
			t.Error("Leaf parent was NOT edge")
		}
		if leaf.suffixLink() != nil {
			t.Error("Leaf suffixLink was NOT nil")
		}
		root.addOutgoingEdgeNode(test.firstValue, edge, leaf)
	}
}
