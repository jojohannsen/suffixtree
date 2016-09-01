package suffixtree

import (
	"testing"

	"github.com/jojohannsen/suffixtree"
)

func TestBuilder(t *testing.T) {
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)
	if ukkonen == nil {
		t.Error("NewBuilder returned nil, want a builder")
	}
	tree := ukkonen.Tree()
	if tree == nil {
		t.Error("Builder has no tree")
	} else {
		root := tree.Root()
		if root == nil {
			t.Error("tree has no root")
		}
		if root.IncomingEdge() != nil {
			t.Error("root has incoming edge")
		}
		if root.suffixLink() != nil {
			t.Error("root has suffix link")
		}
	}
}

func TestBuilderExtend_1(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'

	tree := ukkonen.Tree()
	root := tree.Root()
	for _, test := range tests {
		node := root.NodeFollowing(test.key)
		edge := root.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%s: Node not found", test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%s: got %d outgoing, want %d", test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%s: Edge not found", test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%s: got [%d,%d], want [%d,%d]", test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_2(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 0, 1, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'

	tree := ukkonen.Tree()
	root := tree.Root()
	for _, test := range tests {
		node := root.NodeFollowing(test.key)
		edge := root.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%s: Node not found", test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%s: got %d outgoing, want %d", test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%s: Edge not found", test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%s: got [%d,%d], want [%d,%d]", test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_3(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 0, 1, -1},
		{"root", suffixtree.STKey(rune('s')), 0, 2, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'

	tree := ukkonen.Tree()
	root := tree.Root()
	for _, test := range tests {
		node := root.NodeFollowing(test.key)
		edge := root.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%s: Node not found", test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%s: got %d outgoing, want %d", test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%s: Edge not found", test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%s: got [%d,%d], want [%d,%d]", test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_4(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 0, 1, -1},
		{"root", suffixtree.STKey(rune('s')), 0, 2, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'

	tree := ukkonen.Tree()
	root := tree.Root()
	for _, test := range tests {
		node := root.NodeFollowing(test.key)
		edge := root.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%s: Node not found", test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%s: got %d outgoing, want %d", test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%s: Edge not found", test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%s: got [%d,%d], want [%d,%d]", test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_5(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 0, 1, -1},
		{"root", suffixtree.STKey(rune('s')), 2, 2, 2},
		{"s", suffixtree.STKey(rune('s')), 0, 3, -1},
		{"s", suffixtree.STKey(rune('i')), 0, 4, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'

	tree := ukkonen.Tree()
	root := tree.Root()
	for _, test := range tests {
		baseNode := suffixtree.StrToNode(root, test.title)
		node := baseNode.NodeFollowing(test.key)
		edge := baseNode.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%s: Node not found", test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%s: got %d outgoing, want %d", test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%s: Edge not found", test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%s: got [%d,%d], want [%d,%d]", test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_6(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 0, 1, -1},
		{"root", suffixtree.STKey(rune('s')), 2, 2, 2},
		{"s", suffixtree.STKey(rune('s')), 0, 3, -1},
		{"s", suffixtree.STKey(rune('i')), 0, 4, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'

	tree := ukkonen.Tree()
	root := tree.Root()
	for _, test := range tests {
		baseNode := suffixtree.StrToNode(root, test.title)
		node := baseNode.NodeFollowing(test.key)
		edge := baseNode.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%s: Node not found", test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%s: got %d outgoing, want %d", test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%s: Edge not found", test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%s: got [%d,%d], want [%d,%d]", test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_7(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 0, 1, -1},
		{"root", suffixtree.STKey(rune('s')), 2, 2, 2},
		{"s", suffixtree.STKey(rune('s')), 0, 3, -1},
		{"s", suffixtree.STKey(rune('i')), 0, 4, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'

	tree := ukkonen.Tree()
	root := tree.Root()
	for _, test := range tests {
		baseNode := suffixtree.StrToNode(root, test.title)
		node := baseNode.NodeFollowing(test.key)
		edge := baseNode.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%s: Node not found", test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%s: got %d outgoing, want %d", test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%s: Edge not found", test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%s: got [%d,%d], want [%d,%d]", test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_8(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 0, 1, -1},
		{"root", suffixtree.STKey(rune('s')), 2, 2, 2},
		{"s", suffixtree.STKey(rune('s')), 0, 3, -1},
		{"s", suffixtree.STKey(rune('i')), 0, 4, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'

	tree := ukkonen.Tree()
	root := tree.Root()
	for _, test := range tests {
		baseNode := suffixtree.StrToNode(root, test.title)
		node := baseNode.NodeFollowing(test.key)
		edge := baseNode.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%s: Node not found", test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%s: got %d outgoing, want %d", test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%s: Edge not found", test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%s: got [%d,%d], want [%d,%d]", test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_9(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 2, 1, 1},
		{"root", suffixtree.STKey(rune('s')), 2, 2, 2},
		{"root", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"s", suffixtree.STKey(rune('s')), 2, 3, 4},
		{"s", suffixtree.STKey(rune('i')), 2, 4, 4},
		{"i", suffixtree.STKey(rune('s')), 2, 2, 4},
		{"i", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"s,i", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"s,i", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"s,s", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"s,s", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"i,s", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"i,s", suffixtree.STKey(rune('p')), 0, 8, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 'p'

	tree := ukkonen.Tree()
	root := tree.Root()
	for i, test := range tests {
		baseNode := suffixtree.StrToNode(root, test.title)
		node := baseNode.NodeFollowing(test.key)
		edge := baseNode.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%d,%s: Node not found", i, test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%d,%s: got %d outgoing, want %d", i, test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%d,%s: Edge not found", i, test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%d,%s: got [%d,%d], want [%d,%d]", i, test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_10(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 2, 1, 1},
		{"root", suffixtree.STKey(rune('s')), 2, 2, 2},
		{"root", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"s", suffixtree.STKey(rune('s')), 2, 3, 4},
		{"s", suffixtree.STKey(rune('i')), 2, 4, 4},
		{"i", suffixtree.STKey(rune('s')), 2, 2, 4},
		{"i", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"s,i", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"s,i", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"s,s", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"s,s", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"i,s", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"i,s", suffixtree.STKey(rune('p')), 0, 8, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 'p'
	ukkonen.Extend() // 'p'

	tree := ukkonen.Tree()
	root := tree.Root()
	for i, test := range tests {
		baseNode := suffixtree.StrToNode(root, test.title)
		node := baseNode.NodeFollowing(test.key)
		edge := baseNode.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%d,%s: Node not found", i, test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%d,%s: got %d outgoing, want %d", i, test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%d,%s: Edge not found", i, test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%d,%s: got [%d,%d], want [%d,%d]", i, test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_11(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 2, 1, 1},
		{"root", suffixtree.STKey(rune('s')), 2, 2, 2},
		{"root", suffixtree.STKey(rune('p')), 2, 8, 8},
		{"s", suffixtree.STKey(rune('s')), 2, 3, 4},
		{"s", suffixtree.STKey(rune('i')), 2, 4, 4},
		{"i", suffixtree.STKey(rune('s')), 2, 2, 4},
		{"i", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"p", suffixtree.STKey(rune('p')), 0, 9, -1},
		{"p", suffixtree.STKey(rune('i')), 0, 10, -1},
		{"s,i", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"s,i", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"s,s", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"s,s", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"i,s", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"i,s", suffixtree.STKey(rune('p')), 0, 8, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 'p'
	ukkonen.Extend() // 'p'
	ukkonen.Extend() // 'i'

	tree := ukkonen.Tree()
	root := tree.Root()
	for i, test := range tests {
		baseNode := suffixtree.StrToNode(root, test.title)
		node := baseNode.NodeFollowing(test.key)
		edge := baseNode.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%d,%s: Node not found", i, test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%d,%s: got %d outgoing, want %d", i, test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%d,%s: Edge not found", i, test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%d,%s: got [%d,%d], want [%d,%d]", i, test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}

func TestBuilderExtend_12(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1},
		{"root", suffixtree.STKey(rune('i')), 2, 1, 1},
		{"root", suffixtree.STKey(rune('s')), 2, 2, 2},
		{"root", suffixtree.STKey(rune('p')), 2, 8, 8},
		{"s", suffixtree.STKey(rune('s')), 2, 3, 4},
		{"s", suffixtree.STKey(rune('i')), 2, 4, 4},
		{"i", suffixtree.STKey(rune('s')), 2, 2, 4},
		{"i", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"p", suffixtree.STKey(rune('p')), 0, 9, -1},
		{"p", suffixtree.STKey(rune('i')), 0, 10, -1},
		{"s,i", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"s,i", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"s,s", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"s,s", suffixtree.STKey(rune('p')), 0, 8, -1},
		{"i,s", suffixtree.STKey(rune('s')), 0, 5, -1},
		{"i,s", suffixtree.STKey(rune('p')), 0, 8, -1},
	}
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 'p'
	ukkonen.Extend() // 'p'
	ukkonen.Extend() // 'i'
	if ukkonen.Extend() {
		t.Error("Unexpected extend after reaching end of data")
	}

	tree := ukkonen.Tree()
	root := tree.Root()
	for i, test := range tests {
		baseNode := suffixtree.StrToNode(root, test.title)
		node := baseNode.NodeFollowing(test.key)
		edge := baseNode.EdgeFollowing(test.key)
		if node == nil {
			t.Errorf("%d,%s: Node not found", i, test.title)
		}
		if node.NumberOutgoing() != test.numberOutgoing {
			t.Errorf("%d,%s: got %d outgoing, want %d", i, test.title, node.NumberOutgoing(), test.numberOutgoing)
		}
		if edge == nil {
			t.Errorf("%d,%s: Edge not found", i, test.title)
		}
		if edge.StartOffset != test.incomingEdgeStartOffset || edge.EndOffset != test.incomingEdgeEndOffset {
			t.Errorf("%d,%s: got [%d,%d], want [%d,%d]", i, test.title, edge.StartOffset, edge.EndOffset, test.incomingEdgeStartOffset, test.incomingEdgeEndOffset)
		}
	}
}
