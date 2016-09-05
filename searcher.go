package suffixtree

import "sort"

type Searcher interface {
	Find(sequence []STKey) (suffixOffsets []int32)
}

type searcher struct {
	root       Node
	dataSource DataSource
	traverser  Traverser
}

func NewSearcher(root Node, dataSource DataSource) Searcher {
	return &searcher{root, dataSource, NewTraverser(dataSource)}
}

type int32arr []int32

func (a int32arr) Len() int           { return len(a) }
func (a int32arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int32arr) Less(i, j int) bool { return a[i] < a[j] }

func (s *searcher) Find(sequence []STKey) []int32 {
	result := int32arr{}
	location := NewLocation(s.root)
	for _, val := range sequence {
		if !s.traverser.traverseDownValue(location, val) {
			return result
		}
	}

	result = location.Base.ChildSuffixes(result)
	sort.Sort(result)
	return result
}
