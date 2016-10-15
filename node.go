package suffixtree

import "fmt"

const UnspecifiedOffset int32 = -1
const MoreThanOne = -1

type STKey int

type Node interface {
	// general information about the Node
	String() string
	isRoot() bool
	isInternal() bool
	IsLeaf() bool
	SuffixOffset() int32                         // leaf only
	ChildSuffixes(suffixOffsets []int32) []int32 // all child suffixes
	depth() int32
	Id() int32

	// parent Node and incoming Edge
	parent() Node
	IncomingEdge() *Edge
	setIncoming(node Node, edge *Edge)

	// suffix link
	SuffixLink() Node
	SetSuffixLink(Node)

	// child Nodes and outgoing Edges
	AddOutgoingEdgeNode(key STKey, edge *Edge, node Node)
	outgoingEdgeNode(key STKey) (*Edge, Node)
	addLeafEdgeNode(id int32, key STKey, offset int32) (*Edge, Node)
	EdgeFollowing(key STKey) *Edge
	NodeFollowing(key STKey) Node
	OutgoingNodes() []Node
	outgoingNodeMap() map[STKey]Node
	OutgoingEdgeMap() map[STKey]*Edge
	NumberOutgoing() int
}

func printPathToNode(node Node, dataSource DataSource) {
	result := ""
	for node.parent() != nil {
		result = fmt.Sprintf("%s%s", dataSource.StringFrom(node.IncomingEdge().StartOffset, node.IncomingEdge().EndOffset), result)
		node = node.parent()
	}

	fmt.Println(result)
}

func pathToNode(node Node, dataSource DataSource) string {
	result := ""
	for node.parent() != nil {
		result = fmt.Sprintf("%s%s", dataSource.StringFrom(node.IncomingEdge().StartOffset, node.IncomingEdge().EndOffset), result)
		node = node.parent()
	}
	return result
}

type hasId struct {
	_id int32
}

func (idProvider *hasId) Id() int32 {
	return idProvider._id
}

type idFactory struct {
	_id int32
}

func (factory *idFactory) NextId() int32 {
	factory._id++
	return factory._id
}

func NewNodeIdFactory() (factory *idFactory) {
	return &idFactory{0}
}

// Outgoing edges common to Root and Internal Nodes
type hasOutgoing struct {
	edges map[STKey]*Edge
	nodes map[STKey]Node
}

func (outgoing *hasOutgoing) EdgeFollowing(key STKey) *Edge {
	return outgoing.edges[key]
}

func (outgoing *hasOutgoing) NodeFollowing(key STKey) Node {
	return outgoing.nodes[key]
}

func (outgoing *hasOutgoing) AddOutgoingEdgeNode(key STKey, edge *Edge, node Node) {
	outgoing.edges[key] = edge
	outgoing.nodes[key] = node
}

func (outgoing *hasOutgoing) removeEdgeFollowing(key STKey) {
	delete(outgoing.edges, key)
}

func (outgoing *hasOutgoing) OutgoingNodes() []Node {
	result := []Node{}
	for _, v := range outgoing.nodes {
		result = append(result, v)
	}
	return result
}

func (outgoing *hasOutgoing) outgoingNodeMap() map[STKey]Node {
	return outgoing.nodes
}

func (outgoing *hasOutgoing) OutgoingEdgeMap() map[STKey]*Edge {
	return outgoing.edges
}

func (outgoing *hasOutgoing) outgoingEdgeNode(key STKey) (*Edge, Node) {
	return outgoing.edges[key], outgoing.nodes[key]
}

func (outgoing *hasOutgoing) NumberOutgoing() int {
	return len(outgoing.edges)
}

func (outgoing *hasOutgoing) IsLeaf() bool {
	return false
}

func (outgoing *hasOutgoing) SuffixOffset() int32 {
	panic("no suffix for internal nodes")
}

type noOutgoing struct{}

func (node *noOutgoing) EdgeFollowing(key STKey) *Edge {
	panic("Leaf has no children")
}

func (node *noOutgoing) AddOutgoingEdgeNode(key STKey, edge *Edge, n Node) {
	panic("Leaf has no outgoing edges")
}

func (node *noOutgoing) outgoingEdgeNode(key STKey) (*Edge, Node) {
	panic("Leaf has no outgoing edges or nodes")
}

func (node *noOutgoing) NodeFollowing(key STKey) Node {
	panic("Leaf has no child nodes")
}

func (node *noOutgoing) removeEdgeFollowing(key STKey) {
	panic("Leaf has no children")
}

func (node *noOutgoing) NumberOutgoing() int {
	return 0
}

func (outgoing *noOutgoing) OutgoingNodes() []Node {
	return make([]Node, 0)
}

func (outgoing *noOutgoing) outgoingNodeMap() map[STKey]Node {
	return nil
}

func (outgoing *noOutgoing) OutgoingEdgeMap() map[STKey]*Edge {
	return nil
}

func (outgoing *noOutgoing) IsLeaf() bool {
	return true
}

// Parent edge common to Internal and Leaf Nodes
type hasIncomingEdge struct {
	_parent       Node
	_incomingEdge *Edge
}

func (node *hasIncomingEdge) isRoot() bool {
	return false
}

func (node *hasIncomingEdge) parent() Node {
	return node._parent
}

func (node *hasIncomingEdge) IncomingEdge() *Edge {
	return node._incomingEdge
}

func (node *hasIncomingEdge) setIncoming(parent Node, incomingEdge *Edge) {
	node._parent = parent
	node._incomingEdge = incomingEdge
}

func (node *hasIncomingEdge) depth() int32 {
	return node._incomingEdge.length() + node._parent.depth()
}

type noIncomingEdge struct{}

func (node *noIncomingEdge) isRoot() bool {
	return true
}

func (node *noIncomingEdge) IncomingEdge() *Edge {
	return nil
}

func (node *noIncomingEdge) parent() Node {
	return nil
}

func (node *noIncomingEdge) setIncoming(parent Node, incomingEdge *Edge) {
	panic("Cannot set incoming")
}

func (node *noIncomingEdge) depth() int32 {
	return 0
}

// only Internal Nodes have suffix links
type hasSuffixLink struct {
	_suffixLink Node
}

func (internal *hasSuffixLink) SuffixLink() Node {
	return internal._suffixLink
}

func (internal *hasSuffixLink) SetSuffixLink(slink Node) {
	internal._suffixLink = slink
}

type noSuffixLink struct{}

func (nsl *noSuffixLink) SuffixLink() Node {
	return nil
}

func (nsl *noSuffixLink) SetSuffixLink(node Node) {
	panic("Node does NOT have a suffix link")
}

// Root Node
type rootNode struct {
	hasId
	hasOutgoing
	noIncomingEdge
	noSuffixLink
}

func NewRootNode(id int32) Node {
	return &rootNode{
		hasId{id},
		hasOutgoing{make(map[STKey]*Edge), make(map[STKey]Node)},
		noIncomingEdge{},
		noSuffixLink{}}
}

func (root *rootNode) String() string {
	result := "ROOT("
	for k, v := range root.edges {
		result = fmt.Sprintf("%s%c%s,", result, rune(k), v)
	}
	return fmt.Sprintf("%s)", result)
}

func (root *rootNode) isInternal() bool {
	return false
}

func (root *rootNode) addLeafEdgeNode(id int32, key STKey, offset int32) (*Edge, Node) {
	edge, node := NewLeafEdgeNode(id, root, offset)
	root.AddOutgoingEdgeNode(key, edge, node)
	return edge, node
}

func (root *rootNode) ChildSuffixes(result []int32) []int32 {
	for _, node := range root.OutgoingNodes() {
		result = node.ChildSuffixes(result)
	}
	return result
}

// Internal node
type internalNode struct {
	hasId
	hasOutgoing
	hasIncomingEdge
	hasSuffixLink
}

func NewInternalNode(id int32, parent Node, incoming *Edge) Node {
	return &internalNode{
		hasId{id},
		hasOutgoing{make(map[STKey]*Edge), make(map[STKey]Node)},
		hasIncomingEdge{parent, incoming},
		hasSuffixLink{nil}}
}

func (internal *internalNode) String() string {
	result := fmt.Sprintf("%s Internal(", internal._incomingEdge)
	for k, v := range internal.edges {
		result = fmt.Sprintf("%s%c%s,", result, rune(k), v)
	}
	result = fmt.Sprintf("%s)", result)
	suffixLink := "nil"
	if internal.SuffixLink() != nil {
		suffixLink = fmt.Sprintf("%s", internal.SuffixLink())
	}
	result = fmt.Sprintf("%s, suffixLink=%s", result, suffixLink)
	return result
}

func (internal *internalNode) isInternal() bool {
	return true
}

func (internal *internalNode) addLeafEdgeNode(id int32, key STKey, offset int32) (*Edge, Node) {
	edge, node := NewLeafEdgeNode(id, internal, offset)
	internal.AddOutgoingEdgeNode(key, edge, node)
	return edge, node
}

func (internal *internalNode) ChildSuffixes(result []int32) []int32 {
	for _, node := range internal.OutgoingNodes() {
		result = node.ChildSuffixes(result)
	}
	return result
}

// Leaf node
type leafNode struct {
	hasId
	noOutgoing
	hasIncomingEdge
	noSuffixLink
	_suffixOffset int32
}

func NewLeafEdgeNode(id int32, parent Node, suffix int32) (*Edge, Node) {
	leafEdge := NewLeafEdge(suffix)
	return leafEdge, &leafNode{
		hasId{id},
		noOutgoing{},
		hasIncomingEdge{parent, leafEdge},
		noSuffixLink{}, suffix - parent.depth()}
}

func (leaf *leafNode) String() string {
	return fmt.Sprintf("%s LEAF-%d", leaf._incomingEdge, leaf._suffixOffset)
}

func (leaf *leafNode) isInternal() bool {
	return false
}

func (leaf *leafNode) addLeafEdgeNode(id int32, key STKey, offset int32) (*Edge, Node) {
	panic("Leaf cannot have children")
}

func (leaf *leafNode) SuffixOffset() int32 {
	return leaf._suffixOffset
}

func (leaf *leafNode) ChildSuffixes(result []int32) []int32 {
	return append(result, leaf.SuffixOffset())
}
