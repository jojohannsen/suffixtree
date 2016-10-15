// A suffix tree is built from any data source that has a sequence of values
// that can be mapped to a corresponding sequence of STKey values.
//
// Some trees are built from multiple data sources, so when we store an value offset
// we also need an identifier for the corresponding data source.
package suffixtree

type SuffixTree interface {
	Root() Node
	DataSource() DataSource
}

type suffixTree struct {
	_root       Node
	_dataSource DataSource
}

func NewSuffixTree(root Node, dataSource DataSource) *suffixTree {
	return &suffixTree{root, dataSource}
}

func (st *suffixTree) Root() Node {
	return st._root
}

func (st *suffixTree) DataSource() DataSource {
	return st._dataSource
}