package suffixtree

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func treePrintWithTitle(title string, node Node, location *Location) {
	fmt.Printf("=== %s === %s\n", title, location)
	treePrint(node, 1, "__")
}

func treePrint(node Node, indentLevel int, indentString string) {
	fmt.Printf("%s%s\n", strings.Repeat(indentString, indentLevel), node)
	for k, node := range node.OutgoingNodes() {
		if node == nil {
			fmt.Printf("node is nil, key is %d\n", k)
		} else {
			treePrint(node, indentLevel+1, indentString)
		}
	}
}

func StrToNode(root Node, s string) Node {
	if s == "root" {
		return root
	} else {
		strs := strings.Split(s, ",")
		node := root
		for _, s := range strs {
			node = node.NodeFollowing(STKey(rune(s[0])))
		}
		return node
	}
}

func suffixChildrenIncludesNodeChildren(nodeChildren, suffixChildren []STKey) bool {
	for i := 0; i < len(nodeChildren); i++ {
		found := false
		fmt.Printf("Checking offset %d, value %d\n", i, nodeChildren[i])
		for j := 0; j < len(suffixChildren); j++ {
			fmt.Printf("compare %d to %d\n", nodeChildren[i], suffixChildren[j])
			if nodeChildren[i] == suffixChildren[j] {
				fmt.Printf("FOUND IT!\n")
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

type stkarr []STKey

func (a stkarr) Len() int           { return len(a) }
func (a stkarr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a stkarr) Less(i, j int) bool { return a[i] < a[j] }
func TreeCheck(node Node, dataSource DataSource) {
	edgeChildren := stkarr{}
	nodeChildren := stkarr{}
	for k, _ := range node.outgoingNodeMap() {
		nodeChildren = append(nodeChildren, k)
	}
	for k, _ := range node.OutgoingEdgeMap() {
		edgeChildren = append(edgeChildren, k)
	}
	sort.Sort(nodeChildren)
	sort.Sort(edgeChildren)
	if !reflect.DeepEqual(nodeChildren, edgeChildren) {
		fmt.Printf("Badly constructed node")
		panic("HELP!")
	}
	if node.isInternal() {
		if node.SuffixLink() == nil {
			fmt.Printf("Internal node has no suffix link")
			panic("HELP!")
		}
		suffixNodeChildren := stkarr{}
		for k, _ := range node.SuffixLink().outgoingNodeMap() {
			suffixNodeChildren = append(suffixNodeChildren, k)
		}
		sort.Sort(suffixNodeChildren)
		//if !suffixChildrenIncludesNodeChildren(nodeChildren, suffixNodeChildren) {
		//	fmt.Printf("Suffix doesn't have same children as place linking to it")
		//	fmt.Println("nodeChildren: ", nodeChildren, ", suffixNodeChildren: ", suffixNodeChildren)
		//panic("HELP!")
		//}
		pathToxNode := pathToNode(node, dataSource)
		pathToSuffix := pathToNode(node.SuffixLink(), dataSource)
		if len(pathToxNode) != (len(pathToSuffix) + 1) {
			fmt.Println("Unexpected path lengths, pathToNode=", pathToxNode, ", pathToSuffix=", pathToSuffix)
			panic("HELP!")
		}
	}
	for _, n := range node.OutgoingNodes() {
		TreeCheck(n, dataSource)
	}
}
