package suffixtree

import "fmt"

const UnspecifiedOffset int64 = -1
const MoreThanOne = -1

type STKey int

type Node interface {
	// general information about the Node
	String() string
	isRoot() bool
	isInternal() bool
	IsLeaf() bool
	SuffixOffset() int64 // leaf only
	depth() int64
	Id() int32

	// parent Node and incoming Edge
	parent() Node
	IncomingEdge() *Edge
	setIncoming(node Node, edge *Edge)

	// suffix link
	suffixLink() Node
	setSuffixLink(Node)

	// child Nodes and outgoing Edges
	addOutgoingEdgeNode(key STKey, edge *Edge, node Node)
	outgoingEdgeNode(key STKey) (*Edge, Node)
	addLeafEdgeNode(key STKey, offset int64) (*Edge, Node)
	edgeFollowing(key STKey) *Edge
	NodeFollowing(key STKey) Node
	OutgoingNodes() []Node
	outgoingNodeMap() map[STKey]Node
	OutgoingEdgeMap() map[STKey]*Edge
	numberOutgoing() int
}

func printPathToNode(node Node, dataSource DataSource) {
	result := ""
	for node.parent() != nil {
		result = fmt.Sprintf("%s%s", dataSource.stringFrom(node.IncomingEdge().StartOffset, node.IncomingEdge().EndOffset), result)
		node = node.parent()
	}

	fmt.Println(result)
}

func pathToNode(node Node, dataSource DataSource) string {
	result := ""
	for node.parent() != nil {
		result = fmt.Sprintf("%s%s", dataSource.stringFrom(node.IncomingEdge().StartOffset, node.IncomingEdge().EndOffset), result)
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

func (factory *idFactory) nextId() int32 {
	factory._id++
	return factory._id
}

var nodeIdFactory = &idFactory{0}

// Outgoing edges common to Root and Internal Nodes
type hasOutgoing struct {
	edges map[STKey]*Edge
	nodes map[STKey]Node
}

func (outgoing *hasOutgoing) edgeFollowing(key STKey) *Edge {
	return outgoing.edges[key]
}

func (outgoing *hasOutgoing) NodeFollowing(key STKey) Node {
	return outgoing.nodes[key]
}

func (outgoing *hasOutgoing) addOutgoingEdgeNode(key STKey, edge *Edge, node Node) {
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

func (outgoing *hasOutgoing) numberOutgoing() int {
	return len(outgoing.edges)
}

func (outgoing *hasOutgoing) IsLeaf() bool {
	return false
}

func (outgoing *hasOutgoing) SuffixOffset() int64 {
	panic("no suffix for internal nodes")
}

type noOutgoing struct{}

func (node *noOutgoing) edgeFollowing(key STKey) *Edge {
	panic("Leaf has no children")
}

func (node *noOutgoing) addOutgoingEdgeNode(key STKey, edge *Edge, n Node) {
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

func (node *noOutgoing) numberOutgoing() int {
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

func (node *hasIncomingEdge) depth() int64 {
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

func (node *noIncomingEdge) depth() int64 {
	return 0
}

// only Internal Nodes have suffix links
type hasSuffixLink struct {
	_suffixLink Node
}

func (internal *hasSuffixLink) suffixLink() Node {
	return internal._suffixLink
}

func (internal *hasSuffixLink) setSuffixLink(slink Node) {
	internal._suffixLink = slink
}

type noSuffixLink struct{}

func (nsl *noSuffixLink) suffixLink() Node {
	return nil
}

func (nsl *noSuffixLink) setSuffixLink(node Node) {
	panic("Node does NOT have a suffix link")
}

// Root Node
type rootNode struct {
	hasId
	hasOutgoing
	noIncomingEdge
	noSuffixLink
}

func NewRootNode() Node {
	return &rootNode{
		hasId{nodeIdFactory.nextId()},
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

func (root *rootNode) addLeafEdgeNode(key STKey, offset int64) (*Edge, Node) {
	edge, node := NewLeafEdgeNode(root, offset)
	root.addOutgoingEdgeNode(key, edge, node)
	return edge, node
}

// Internal node
type internalNode struct {
	hasId
	hasOutgoing
	hasIncomingEdge
	hasSuffixLink
}

func NewInternalNode(parent Node, incoming *Edge) Node {
	return &internalNode{
		hasId{nodeIdFactory.nextId()},
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
	if internal.suffixLink() != nil {
		suffixLink = fmt.Sprintf("%s", internal.suffixLink())
	}
	result = fmt.Sprintf("%s, suffixLink=%s", result, suffixLink)
	return result
}

func (internal *internalNode) isInternal() bool {
	return true
}

func (internal *internalNode) addLeafEdgeNode(key STKey, offset int64) (*Edge, Node) {
	edge, node := NewLeafEdgeNode(internal, offset)
	internal.addOutgoingEdgeNode(key, edge, node)
	return edge, node
}

// Leaf node
type leafNode struct {
	hasId
	noOutgoing
	hasIncomingEdge
	noSuffixLink
	_suffixOffset int64
}

func NewLeafEdgeNode(parent Node, suffix int64) (*Edge, Node) {
	leafEdge := NewLeafEdge(suffix)
	return leafEdge, &leafNode{
		hasId{nodeIdFactory.nextId()},
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

func (leaf *leafNode) addLeafEdgeNode(key STKey, offset int64) (*Edge, Node) {
	panic("Leaf cannot have children")
}

func (leaf *leafNode) SuffixOffset() int64 {
	return leaf._suffixOffset
}