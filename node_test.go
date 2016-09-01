package suffixtree

import (
	"strconv"
	"testing"

	"github.com/jojohannsen/suffixtree"
)

func TestRootNode(t *testing.T) {
	var root suffixtree.Node = suffixtree.NewRootNode()
	if root == nil {
		t.Error("NewRootNode returned nil")
	}
	if root.IncomingEdge() != nil {
		t.Error("root has unexpected parent")
	}
	if root.EdgeFollowing(0) != nil {
		t.Error("unexpected children")
	}
	if root.SuffixLink() != nil {
		t.Error("Unexpected suffix link")
	}
}

func TestRootOutgoingEdge(t *testing.T) {
	root := suffixtree.NewRootNode()
	edge := suffixtree.NewEdge(0, 3)
	root.AddOutgoingEdgeNode(1000, edge, nil)
	testEdge := root.EdgeFollowing(1000)
	if testEdge != edge {
		t.Error("Did not get edge just added")
	}
	testEdge = root.EdgeFollowing(1001)
	if testEdge != nil {
		t.Error("Found edge where none expected")
	}
}

func TestRootPanic1(t *testing.T) {
	defer func() {
		recover()
	}()
	root := suffixtree.NewRootNode()
	root.SetSuffixLink(nil)
	t.Errorf("The code did not panic")
}

func TestRootChildren(t *testing.T) {
	var rootTests = []struct {
		startOffset int64
		endOffset   int64
		firstValue  suffixtree.STKey
	}{
		{0, 3, 1001},
		{1, 3, 1002},
		{2, 3, 1003},
		{3, 3, 1004},
	}
	root := suffixtree.NewRootNode()
	for _, test := range rootTests {
		edge := suffixtree.NewEdge(test.startOffset, test.endOffset)
		root.AddOutgoingEdgeNode(test.firstValue, edge, nil)
	}
	for _, test := range rootTests {
		testEdge := root.EdgeFollowing(test.firstValue)
		if testEdge.StartOffset != test.startOffset {
			t.Error("Edge start offset, got " + strconv.FormatInt(testEdge.StartOffset, 10) + ", want " + strconv.FormatInt(test.startOffset, 10))
		}
		if testEdge.EndOffset != test.endOffset {
			t.Error("Edge end offset, got " + strconv.FormatInt(testEdge.EndOffset, 10) + ", want " + strconv.FormatInt(test.endOffset, 10))
		}
	}
	testEdge := root.EdgeFollowing(111111)
	if testEdge != nil {
		t.Error("Found edge where none expected")
	}
}

func TestInternalChildren(t *testing.T) {
	var internalTests = []struct {
		startOffset int64
		endOffset   int64
		firstValue  suffixtree.STKey
	}{
		{0, 3, 1001},
		{1, 3, 1002},
		{2, 3, 1003},
		{3, 3, 1004},
	}
	root := suffixtree.NewRootNode()
	for _, test := range internalTests {
		edge := suffixtree.NewEdge(test.startOffset, test.endOffset)
		internal := suffixtree.NewInternalNode(root, edge)
		if internal.IncomingEdge() != edge {
			t.Error("Internal parent not found")
		}
		internalEdge := suffixtree.NewEdge(test.startOffset, test.endOffset)
		root.AddOutgoingEdgeNode(test.firstValue, edge, nil)
		internal.AddOutgoingEdgeNode(2000, internalEdge, nil)
		if internal.SuffixLink() != nil {
			t.Error("new Internal Node had a suffix link")
		}
		if internal.EdgeFollowing(2000) != internalEdge {
			t.Error("internal edge follow failed")
		}
		internal.SetSuffixLink(root)
		if internal.SuffixLink() != root {
			t.Error("InternalNode SetSuffixLink failed")
		}
	}
	for _, test := range internalTests {
		testEdge := root.EdgeFollowing(test.firstValue)
		if testEdge.StartOffset != test.startOffset {
			t.Error("Edge start offset, got " + strconv.FormatInt(testEdge.StartOffset, 10) + ", want " + strconv.FormatInt(test.startOffset, 10))
		}
		if testEdge.EndOffset != test.endOffset {
			t.Error("Edge end offset, got " + strconv.FormatInt(testEdge.EndOffset, 10) + ", want " + strconv.FormatInt(test.endOffset, 10))
		}
	}
	testEdge := root.EdgeFollowing(111111)
	if testEdge != nil {
		t.Error("Found edge where none expected")
	}
}

func TestLeafOutgoingFollowing(t *testing.T) {
	defer func() {
		recover()
	}()

	root := suffixtree.NewRootNode()
	edge, leaf := suffixtree.NewLeafEdgeNode(root, 0)
	leaf.AddOutgoingEdgeNode(1000, edge, leaf)
	t.Errorf("The code did not panic")
}

func TestLeafEdgeFollowingPanic(t *testing.T) {
	defer func() {
		recover()
	}()

	root := suffixtree.NewRootNode()
	_, leaf := suffixtree.NewLeafEdgeNode(root, 0)
	leaf.EdgeFollowing(1000)
	t.Errorf("The code did not panic")
}

func TestLeafSuffixLinkPanic(t *testing.T) {
	defer func() {
		recover()
	}()

	root := suffixtree.NewRootNode()
	_, leaf := suffixtree.NewLeafEdgeNode(root, 0)
	leaf.SetSuffixLink(nil)
	t.Errorf("The code did not panic")
}

func TestRootLeaf(t *testing.T) {
	var rootLeafTests = []struct {
		startOffset int64
		endOffset   int64
		firstValue  suffixtree.STKey
	}{
		{0, -1, 1001},
		{1, -1, 1002},
		{2, -1, 1003},
		{3, -1, 1004},
	}
	root := suffixtree.NewRootNode()
	for _, test := range rootLeafTests {
		edge := suffixtree.NewEdge(test.startOffset, test.endOffset)
		edge, leaf := suffixtree.NewLeafEdgeNode(root, test.startOffset)
		if leaf.IncomingEdge() != edge {
			t.Error("Leaf parent was NOT edge")
		}
		if leaf.SuffixLink() != nil {
			t.Error("Leaf suffixLink was NOT nil")
		}
		root.AddOutgoingEdgeNode(test.firstValue, edge, leaf)
	}
}
