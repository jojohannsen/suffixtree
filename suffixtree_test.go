package suffixtree

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"github.com/jojohannsen/suffixtree"
)

func TestX(t *testing.T) {
	tests := []struct {
		title                   string
		key                     suffixtree.STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
		suffixOffset            int64
	}{
		{"root", suffixtree.STKey(rune('m')), 0, 0, -1, 0},
		{"root", suffixtree.STKey(rune('i')), 0, 1, -1, 1},
		{"root", suffixtree.STKey(rune('s')), 2, 2, 2, -1},
		{"root", suffixtree.STKey(rune('$')), 0, 4, -1, 4},
		{"s", suffixtree.STKey(rune('s')), 0, 3, -1, 2},
		{"s", suffixtree.STKey(rune('$')), 0, 4, -1, 3},
	}
	dataSource := suffixtree.NewStringDataSource("miss$")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	tree := ukkonen.Tree()
	ukkonen.Extend()
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
		if test.suffixOffset != -1 {
			if node.SuffixOffset() != test.suffixOffset {
				t.Errorf("%s: suffixOffset, got %d, want %d", test.title, node.SuffixOffset(), test.suffixOffset)
			}
		}
	}
}

func TestSuffixTree(t *testing.T) {
	dataSource := suffixtree.NewStringDataSource("mississippi$")
	ukkonen := suffixtree.NewUkkonen(dataSource)
	ukkonen.DrainDataSource()
	st := ukkonen.Tree()
	searcher := suffixtree.NewSearcher(st.Root(), dataSource)
	m := []suffixtree.STKey{suffixtree.STKey(rune('m'))}
	result := searcher.Find(m)
	if len(result) != 1 {
		t.Error("Did not find 'm'")
	}
	expectedM := []int64{0}
	if !reflect.DeepEqual(result, expectedM) {
		t.Errorf("Find failed, got %s, want %s", result, expectedM)
	}
	i := []suffixtree.STKey{suffixtree.STKey(rune('i'))}
	result = searcher.Find(i)
	expectedI := []int64{1, 4, 7, 10}
	if !reflect.DeepEqual(result, expectedI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedI)
	}
	s := []suffixtree.STKey{suffixtree.STKey(rune('s'))}
	result = searcher.Find(s)
	expectedS := []int64{2, 3, 5, 6}
	if !reflect.DeepEqual(result, expectedS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedS)
	}
	p := []suffixtree.STKey{suffixtree.STKey(rune('p'))}
	result = searcher.Find(p)
	expectedP := []int64{8, 9}
	if !reflect.DeepEqual(result, expectedP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedP)
	}
	dollar := []suffixtree.STKey{suffixtree.STKey(rune('$'))}
	result = searcher.Find(dollar)
	expectedDollar := []int64{11}
	if !reflect.DeepEqual(result, expectedDollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedDollar)
	}
}

func TestSuffixTree2(t *testing.T) {
	dataSource := suffixtree.NewStringDataSource("mississippi$")
	ukkonen := suffixtree.NewUkkonen(dataSource)
	ukkonen.DrainDataSource()
	st := ukkonen.Tree()
	searcher := suffixtree.NewSearcher(st.Root(), dataSource)
	mi := []suffixtree.STKey{suffixtree.STKey(rune('m')), suffixtree.STKey(rune('i'))}
	result := searcher.Find(mi)
	if len(result) != 1 {
		t.Error("Did not find 'mi'")
	}
	expectedMI := []int64{0}
	if !reflect.DeepEqual(result, expectedMI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedMI)
	}
	is := []suffixtree.STKey{suffixtree.STKey(rune('i')), suffixtree.STKey(rune('s'))}
	result = searcher.Find(is)
	expectedIS := []int64{1, 4}
	if !reflect.DeepEqual(result, expectedIS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIS)
	}
	ip := []suffixtree.STKey{suffixtree.STKey(rune('i')), suffixtree.STKey(rune('p'))}
	result = searcher.Find(ip)
	expectedIP := []int64{7}
	if !reflect.DeepEqual(result, expectedIP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIP)
	}
	idollar := []suffixtree.STKey{suffixtree.STKey(rune('i')), suffixtree.STKey(rune('$'))}
	result = searcher.Find(idollar)
	expectedIdollar := []int64{10}
	if !reflect.DeepEqual(result, expectedIdollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIdollar)
	}
	ss := []suffixtree.STKey{suffixtree.STKey(rune('s')), suffixtree.STKey(rune('s'))}
	result = searcher.Find(ss)
	expectedSS := []int64{2, 5}
	if !reflect.DeepEqual(result, expectedSS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSS)
	}
	si := []suffixtree.STKey{suffixtree.STKey(rune('s')), suffixtree.STKey(rune('i'))}
	result = searcher.Find(si)
	expectedSI := []int64{3, 6}
	if !reflect.DeepEqual(result, expectedSI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSI)
	}

	pp := []suffixtree.STKey{suffixtree.STKey(rune('p')), suffixtree.STKey(rune('p'))}
	result = searcher.Find(pp)
	expectedPP := []int64{8}
	if !reflect.DeepEqual(result, expectedPP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedPP)
	}
	pi := []suffixtree.STKey{suffixtree.STKey(rune('p')), suffixtree.STKey(rune('i'))}
	result = searcher.Find(pi)
	expectedPI := []int64{9}
	if !reflect.DeepEqual(result, expectedPI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedPI)
	}
	dollar := []suffixtree.STKey{suffixtree.STKey(rune('$'))}
	result = searcher.Find(dollar)
	expectedDollar := []int64{11}
	if !reflect.DeepEqual(result, expectedDollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedDollar)
	}
}

func TestSuffixTree3(t *testing.T) {
	dataSource := suffixtree.NewStringDataSource("mississippi$")
	ukkonen := suffixtree.NewUkkonen(dataSource)
	ukkonen.DrainDataSource()
	st := ukkonen.Tree()
	searcher := suffixtree.NewSearcher(st.Root(), dataSource)
	mis := []suffixtree.STKey{suffixtree.STKey(rune('m')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('s'))}
	result := searcher.Find(mis)
	if len(result) != 1 {
		t.Error("Did not find 'mis'")
	}
	expectedMIS := []int64{0}
	if !reflect.DeepEqual(result, expectedMIS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedMIS)
	}
	iss := []suffixtree.STKey{suffixtree.STKey(rune('i')), suffixtree.STKey(rune('s')), suffixtree.STKey(rune('s'))}
	result = searcher.Find(iss)
	expectedISS := []int64{1, 4}
	if !reflect.DeepEqual(result, expectedISS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedISS)
	}
	ipp := []suffixtree.STKey{suffixtree.STKey(rune('i')), suffixtree.STKey(rune('p')), suffixtree.STKey(rune('p'))}
	result = searcher.Find(ipp)
	expectedIPP := []int64{7}
	if !reflect.DeepEqual(result, expectedIPP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIPP)
	}
	ssi := []suffixtree.STKey{suffixtree.STKey(rune('s')), suffixtree.STKey(rune('s')), suffixtree.STKey(rune('i'))}
	result = searcher.Find(ssi)
	expectedSSI := []int64{2, 5}
	if !reflect.DeepEqual(result, expectedSSI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSSI)
	}
	sis := []suffixtree.STKey{suffixtree.STKey(rune('s')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('s'))}
	result = searcher.Find(sis)
	expectedSIS := []int64{3}
	if !reflect.DeepEqual(result, expectedSIS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSIS)
	}
	sip := []suffixtree.STKey{suffixtree.STKey(rune('s')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('p'))}
	result = searcher.Find(sip)
	expectedSIP := []int64{6}
	if !reflect.DeepEqual(result, expectedSIP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSIP)
	}

	ppi := []suffixtree.STKey{suffixtree.STKey(rune('p')), suffixtree.STKey(rune('p')), suffixtree.STKey(rune('i'))}
	result = searcher.Find(ppi)
	expectedPPI := []int64{8}
	if !reflect.DeepEqual(result, expectedPPI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedPPI)
	}
	piDollar := []suffixtree.STKey{suffixtree.STKey(rune('p')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('$'))}
	result = searcher.Find(piDollar)
	expectedPIdollar := []int64{9}
	if !reflect.DeepEqual(result, expectedPIdollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedPIdollar)
	}
	dollar := []suffixtree.STKey{suffixtree.STKey(rune('$'))}
	result = searcher.Find(dollar)
	expectedDollar := []int64{11}
	if !reflect.DeepEqual(result, expectedDollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedDollar)
	}
}

func TestSuffixTree4(t *testing.T) {
	dataSource := suffixtree.NewStringDataSource("mississippi$")
	ukkonen := suffixtree.NewUkkonen(dataSource)
	ukkonen.DrainDataSource()
	st := ukkonen.Tree()
	searcher := suffixtree.NewSearcher(st.Root(), dataSource)
	miss := []suffixtree.STKey{suffixtree.STKey(rune('m')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('s')), suffixtree.STKey(rune('s'))}
	result := searcher.Find(miss)
	if len(result) != 1 {
		t.Error("Did not find 'miss'")
	}
	expectedMISS := []int64{0}
	if !reflect.DeepEqual(result, expectedMISS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedMISS)
	}
	issi := []suffixtree.STKey{suffixtree.STKey(rune('i')), suffixtree.STKey(rune('s')), suffixtree.STKey(rune('s')), suffixtree.STKey(rune('i'))}
	result = searcher.Find(issi)
	expectedISSI := []int64{1, 4}
	if !reflect.DeepEqual(result, expectedISSI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedISSI)
	}
	ippi := []suffixtree.STKey{suffixtree.STKey(rune('i')), suffixtree.STKey(rune('p')), suffixtree.STKey(rune('p')), suffixtree.STKey(rune('i'))}
	result = searcher.Find(ippi)
	expectedIPPI := []int64{7}
	if !reflect.DeepEqual(result, expectedIPPI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIPPI)
	}
	ssis := []suffixtree.STKey{suffixtree.STKey(rune('s')), suffixtree.STKey(rune('s')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('s'))}
	result = searcher.Find(ssis)
	expectedSSIS := []int64{2}
	if !reflect.DeepEqual(result, expectedSSIS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSSIS)
	}
	ssip := []suffixtree.STKey{suffixtree.STKey(rune('s')), suffixtree.STKey(rune('s')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('p'))}
	result = searcher.Find(ssip)
	expectedSSIP := []int64{5}
	if !reflect.DeepEqual(result, expectedSSIP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSSIP)
	}
	siss := []suffixtree.STKey{suffixtree.STKey(rune('s')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('s')), suffixtree.STKey(rune('s'))}
	result = searcher.Find(siss)
	expectedSISS := []int64{3}
	if !reflect.DeepEqual(result, expectedSISS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSISS)
	}
	sipp := []suffixtree.STKey{suffixtree.STKey(rune('s')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('p')), suffixtree.STKey(rune('p'))}
	result = searcher.Find(sipp)
	expectedSIPP := []int64{6}
	if !reflect.DeepEqual(result, expectedSIPP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSIPP)
	}

	ppiDollar := []suffixtree.STKey{suffixtree.STKey(rune('p')), suffixtree.STKey(rune('p')), suffixtree.STKey(rune('i')), suffixtree.STKey(rune('$'))}
	result = searcher.Find(ppiDollar)
	expectedPPIdollar := []int64{8}
	if !reflect.DeepEqual(result, expectedPPIdollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedPPIdollar)
	}
}

var letterRunes = []rune("ABCD") //"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) []rune {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return b
}

func sliceappend(old, new []rune) []rune {
	newlen := len(old) + len(new)
	newcap := newlen * 2
	newSlice := make([]rune, newlen, newcap)
	copy(newSlice, old)
	copy(newSlice[len(old):], new)
	return newSlice
}

func convertToSTKey(runes []rune) []suffixtree.STKey {
	result := []suffixtree.STKey{}
	for i := 0; i < len(runes); i++ {
		result = append(result, suffixtree.STKey(runes[i]))
	}
	return result
}

func TestRandomString(t *testing.T) {
	rand.Seed(1)

	for i := 0; i < 100; i++ {
		s := []rune{}
		s = sliceappend(s, RandStringRunes(i))
		fmt.Printf("s=%s\n", string(s))
		/*s := []rune("0eeijeieijj")
		s = sliceappend(s, RandStringRunes(990))
		s = sliceappend(s, []rune("mississippi"))
		//fmt.Println("Len is ", len(s))
		s = sliceappend(s, RandStringRunes(989))
		s = sliceappend(s, []rune("mississippi"))
		s = sliceappend(s, RandStringRunes(100))*/

		dataSource := suffixtree.NewRuneDataSource(s)
		ukkonen := suffixtree.NewUkkonen(dataSource)
		ukkonen.DrainDataSource()
		ukkonen.Finish()
		st := ukkonen.Tree()
		//treePrintWithTitle(fmt.Sprintf("Test '%s'", string(s)), st.Root(), ukkonen.Location())
		suffixtree.TreeCheck(st.Root(), dataSource)
		for i := int64(0); i < int64(len(s)); i++ {
			searcher := suffixtree.NewSearcher(st.Root(), dataSource)
			searchFor := s[i:len(s)]
			fmt.Printf("search for suffix %s\n", string(searchFor))
			result := searcher.Find(convertToSTKey(searchFor))
			if len(result) == 0 {
				fmt.Printf("Did not find %s at all\n", string(searchFor))
			} else {
				foundSuffix := false
				for j := 0; j < len(result); j++ {
					if result[j] == i {
						foundSuffix = true
					}
				}
				if !foundSuffix {
					fmt.Printf("Did not find suffix %s at expected offset %d\n", string(searchFor), i)
					for j := 0; j < len(result); j++ {
						fmt.Printf("  found at %d\n", result[j])
					}
				}
			}
		}
		//treePrintWithTitle("test", st.Root(), ukkonen.Location())
		/*searcher := NewSearcher(st.Root(), dataSource)
		searchFor := []suffixtree.STKey{}
		searchFor = append(searchFor, suffixtree.STKey(rune('B')))
		result := searcher.find(searchFor)
		if len(result) == 0 {
			t.Errorf("Expected >1 result, got %d", len(result))
		} else if result[0] != 0 {
			t.Errorf("Expected offset 0, got %d", result[0])
		}*/

	}
}
