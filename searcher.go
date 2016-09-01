package suffixtree

import "sort"

type Searcher interface {
	find(sequence []STKey) (suffixOffsets []int64)
}

type searcher struct {
	root Node
	dataSource DataSource
	traverser Traverser
}

func NewSearcher(root Node, dataSource DataSource) Searcher {
	return &searcher{root, dataSource, NewTraverser(dataSource)}
}

func (s *searcher) collectSuffixes(node Node, result []int64) []int64 {
	if node.IsLeaf() {
		result = append(result, node.SuffixOffset())
	} else {
		for _,node := range node.OutgoingNodes() {
			result = s.collectSuffixes(node, result)
		}
	}
	return result
}

type int64arr []int64
func (a int64arr) Len() int { return len(a) }
func (a int64arr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

func (s *searcher) find(sequence[] STKey) ([]int64) {
	result := int64arr{}
	location := NewLocation(s.root)
	for _,val := range sequence {
		if !s.traverser.traverseDownValue(location, val) {
			return result
		}	
	}

	result = s.collectSuffixes(location.Base, result)
	sort.Sort(result)
	return result
}
