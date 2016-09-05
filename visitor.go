package suffixtree

import "fmt"

type VisitorTraverser interface {
	Traverse(node Node)
}

type Visitor interface {
	PreVisit(node Node) bool
	Visit(node Node) bool
	PostVisit(node Node) bool
	Done() bool
}

type hasVisitor struct {
	visitor Visitor
}

// BFS
type BFS struct {
	hasVisitor
	queue []Node
}

func NewBFS(visitor Visitor) *BFS {
	return &BFS{hasVisitor{visitor}, make([]Node, 0)}
}

func (bfs *BFS) Traverse(node Node) {
	bfs.queue = append(bfs.queue, node)
	for (len(bfs.queue) > 0) && !bfs.visitor.Done() {
		node = bfs.queue[0]
		bfs.queue = bfs.queue[1:]
		if bfs.visitor.PreVisit(node) {
			for _, child := range node.OutgoingNodes() {
				bfs.queue = append(bfs.queue, child)
			}
			if bfs.visitor.Visit(node) {
				bfs.visitor.PostVisit(node)
			}
		}
	}
}

// DFS
type DFS struct {
	hasVisitor
}

func NewDFS(visitor Visitor) *DFS {
	return &DFS{hasVisitor{visitor}}
}

func (dfs *DFS) Traverse(node Node) {
	if dfs.visitor.PreVisit(node) {
		if dfs.visitor.Visit(node) {
			for _, child := range node.OutgoingNodes() {
				dfs.Traverse(child)
			}
		}
		dfs.visitor.PostVisit(node)
	}
}

// general descriptive visitor behaviors
type noPostVisit struct{}

func (npv *noPostVisit) PostVisit(node Node) bool {
	return true
}

type noDone struct{}

func (nd *noDone) Done() bool {
	return false
}

// Visitor Implementations

// Suffix Link Printer prints all suffix links for a tree
type SuffixLinkPrinter struct {
	noPostVisit
	noDone
}

func NewSuffixLinkPrinter() *SuffixLinkPrinter {
	return &SuffixLinkPrinter{}
}

func (slp *SuffixLinkPrinter) PreVisit(node Node) bool {
	return true
}

func (slp *SuffixLinkPrinter) Visit(node Node) bool {
	if node.isInternal() {
		sl := node.SuffixLink()
		if sl == nil {
			fmt.Printf("(%d->NIL)", node.Id())
		} else {
			fmt.Printf("(%d->%d)", node.Id(), node.SuffixLink().Id())
			fmt.Printf(" depth %d, %d\n", node.depth(), node.SuffixLink().depth())
		}
		return true
	} else {
		return false
	}
}

// Depth visitor sends out sets of suffixes each time specified depth is reached
type DepthVisitor struct {
	noPostVisit
	noDone
	maxDepth int32
	outChan chan<- []int32
}

func NewDepthVisitor(depth int32, outChan chan<- []int32) *DepthVisitor {
	dv := &DepthVisitor{}
	dv.maxDepth = depth
	dv.outChan = outChan
	return dv
}

func (dv *DepthVisitor) PreVisit(node Node) bool {
	if node.IsLeaf() || (node.depth() >= dv.maxDepth) {
		dv.outChan <- node.ChildSuffixes([]int32{})
		return false
	}
	return true
}

func (dv *DepthVisitor) Visit(node Node) bool {
	return true
}


