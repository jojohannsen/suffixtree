package suffixtree

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestX(t *testing.T) {
	tests := []struct {
		title                   string
		key                     STKey
		numberOutgoing          int
		incomingEdgeStartOffset int64
		incomingEdgeEndOffset   int64
		suffixOffset            int64
	}{
		{"root", STKey(rune('m')), 0, 0, -1, 0},
		{"root", STKey(rune('i')), 0, 1, -1, 1},
		{"root", STKey(rune('s')), 2, 2, 2, -1},
		{"root", STKey(rune('$')), 0, 4, -1, 4},
		{"s", STKey(rune('s')), 0, 3, -1, 2},
		{"s", STKey(rune('$')), 0, 4, -1, 3},
	}
	dataSource := NewStringDataSource("miss$")
	ukkonen := NewUkkonen(dataSource)

	ukkonen.Extend() // 'm'
	ukkonen.Extend() // 'i'
	ukkonen.Extend() // 's'
	ukkonen.Extend() // 's'
	tree := ukkonen.Tree()
	ukkonen.Extend()
	root := tree.Root()

	for _, test := range tests {
		baseNode := strToNode(root, test.title)
		node := baseNode.NodeFollowing(test.key)
		edge := baseNode.edgeFollowing(test.key)
		if node == nil {
			t.Errorf("%s: Node not found", test.title)
		}
		if node.numberOutgoing() != test.numberOutgoing {
			t.Errorf("%s: got %d outgoing, want %d", test.title, node.numberOutgoing(), test.numberOutgoing)
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
	dataSource := NewStringDataSource("mississippi$")
	ukkonen := NewUkkonen(dataSource)
	ukkonen.drainDataSource()
	st := ukkonen.Tree()
	searcher := NewSearcher(st.Root(), dataSource)
	m := []STKey{STKey(rune('m'))}
	result := searcher.find(m)
	if len(result) != 1 {
		t.Error("Did not find 'm'")
	}
	expectedM := []int64{0}
	if !reflect.DeepEqual(result, expectedM) {
		t.Errorf("Find failed, got %s, want %s", result, expectedM)
	}
	i := []STKey{STKey(rune('i'))}
	result = searcher.find(i)
	expectedI := []int64{1, 4, 7, 10}
	if !reflect.DeepEqual(result, expectedI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedI)
	}
	s := []STKey{STKey(rune('s'))}
	result = searcher.find(s)
	expectedS := []int64{2, 3, 5, 6}
	if !reflect.DeepEqual(result, expectedS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedS)
	}
	p := []STKey{STKey(rune('p'))}
	result = searcher.find(p)
	expectedP := []int64{8, 9}
	if !reflect.DeepEqual(result, expectedP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedP)
	}
	dollar := []STKey{STKey(rune('$'))}
	result = searcher.find(dollar)
	expectedDollar := []int64{11}
	if !reflect.DeepEqual(result, expectedDollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedDollar)
	}
}

func TestSuffixTree2(t *testing.T) {
	dataSource := NewStringDataSource("mississippi$")
	ukkonen := NewUkkonen(dataSource)
	ukkonen.drainDataSource()
	st := ukkonen.Tree()
	searcher := NewSearcher(st.Root(), dataSource)
	mi := []STKey{STKey(rune('m')), STKey(rune('i'))}
	result := searcher.find(mi)
	if len(result) != 1 {
		t.Error("Did not find 'mi'")
	}
	expectedMI := []int64{0}
	if !reflect.DeepEqual(result, expectedMI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedMI)
	}
	is := []STKey{STKey(rune('i')), STKey(rune('s'))}
	result = searcher.find(is)
	expectedIS := []int64{1, 4}
	if !reflect.DeepEqual(result, expectedIS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIS)
	}
	ip := []STKey{STKey(rune('i')), STKey(rune('p'))}
	result = searcher.find(ip)
	expectedIP := []int64{7}
	if !reflect.DeepEqual(result, expectedIP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIP)
	}
	idollar := []STKey{STKey(rune('i')), STKey(rune('$'))}
	result = searcher.find(idollar)
	expectedIdollar := []int64{10}
	if !reflect.DeepEqual(result, expectedIdollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIdollar)
	}
	ss := []STKey{STKey(rune('s')), STKey(rune('s'))}
	result = searcher.find(ss)
	expectedSS := []int64{2, 5}
	if !reflect.DeepEqual(result, expectedSS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSS)
	}
	si := []STKey{STKey(rune('s')), STKey(rune('i'))}
	result = searcher.find(si)
	expectedSI := []int64{3, 6}
	if !reflect.DeepEqual(result, expectedSI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSI)
	}

	pp := []STKey{STKey(rune('p')), STKey(rune('p'))}
	result = searcher.find(pp)
	expectedPP := []int64{8}
	if !reflect.DeepEqual(result, expectedPP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedPP)
	}
	pi := []STKey{STKey(rune('p')), STKey(rune('i'))}
	result = searcher.find(pi)
	expectedPI := []int64{9}
	if !reflect.DeepEqual(result, expectedPI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedPI)
	}
	dollar := []STKey{STKey(rune('$'))}
	result = searcher.find(dollar)
	expectedDollar := []int64{11}
	if !reflect.DeepEqual(result, expectedDollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedDollar)
	}
}

func TestSuffixTree3(t *testing.T) {
	dataSource := NewStringDataSource("mississippi$")
	ukkonen := NewUkkonen(dataSource)
	ukkonen.drainDataSource()
	st := ukkonen.Tree()
	searcher := NewSearcher(st.Root(), dataSource)
	mis := []STKey{STKey(rune('m')), STKey(rune('i')), STKey(rune('s'))}
	result := searcher.find(mis)
	if len(result) != 1 {
		t.Error("Did not find 'mis'")
	}
	expectedMIS := []int64{0}
	if !reflect.DeepEqual(result, expectedMIS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedMIS)
	}
	iss := []STKey{STKey(rune('i')), STKey(rune('s')), STKey(rune('s'))}
	result = searcher.find(iss)
	expectedISS := []int64{1, 4}
	if !reflect.DeepEqual(result, expectedISS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedISS)
	}
	ipp := []STKey{STKey(rune('i')), STKey(rune('p')), STKey(rune('p'))}
	result = searcher.find(ipp)
	expectedIPP := []int64{7}
	if !reflect.DeepEqual(result, expectedIPP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIPP)
	}
	ssi := []STKey{STKey(rune('s')), STKey(rune('s')), STKey(rune('i'))}
	result = searcher.find(ssi)
	expectedSSI := []int64{2, 5}
	if !reflect.DeepEqual(result, expectedSSI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSSI)
	}
	sis := []STKey{STKey(rune('s')), STKey(rune('i')), STKey(rune('s'))}
	result = searcher.find(sis)
	expectedSIS := []int64{3}
	if !reflect.DeepEqual(result, expectedSIS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSIS)
	}
	sip := []STKey{STKey(rune('s')), STKey(rune('i')), STKey(rune('p'))}
	result = searcher.find(sip)
	expectedSIP := []int64{6}
	if !reflect.DeepEqual(result, expectedSIP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSIP)
	}

	ppi := []STKey{STKey(rune('p')), STKey(rune('p')), STKey(rune('i'))}
	result = searcher.find(ppi)
	expectedPPI := []int64{8}
	if !reflect.DeepEqual(result, expectedPPI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedPPI)
	}
	piDollar := []STKey{STKey(rune('p')), STKey(rune('i')), STKey(rune('$'))}
	result = searcher.find(piDollar)
	expectedPIdollar := []int64{9}
	if !reflect.DeepEqual(result, expectedPIdollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedPIdollar)
	}
	dollar := []STKey{STKey(rune('$'))}
	result = searcher.find(dollar)
	expectedDollar := []int64{11}
	if !reflect.DeepEqual(result, expectedDollar) {
		t.Errorf("Find failed, got %s, want %s", result, expectedDollar)
	}
}

func TestSuffixTree4(t *testing.T) {
	dataSource := NewStringDataSource("mississippi$")
	ukkonen := NewUkkonen(dataSource)
	ukkonen.drainDataSource()
	st := ukkonen.Tree()
	searcher := NewSearcher(st.Root(), dataSource)
	miss := []STKey{STKey(rune('m')), STKey(rune('i')), STKey(rune('s')), STKey(rune('s'))}
	result := searcher.find(miss)
	if len(result) != 1 {
		t.Error("Did not find 'miss'")
	}
	expectedMISS := []int64{0}
	if !reflect.DeepEqual(result, expectedMISS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedMISS)
	}
	issi := []STKey{STKey(rune('i')), STKey(rune('s')), STKey(rune('s')), STKey(rune('i'))}
	result = searcher.find(issi)
	expectedISSI := []int64{1, 4}
	if !reflect.DeepEqual(result, expectedISSI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedISSI)
	}
	ippi := []STKey{STKey(rune('i')), STKey(rune('p')), STKey(rune('p')), STKey(rune('i'))}
	result = searcher.find(ippi)
	expectedIPPI := []int64{7}
	if !reflect.DeepEqual(result, expectedIPPI) {
		t.Errorf("Find failed, got %s, want %s", result, expectedIPPI)
	}
	ssis := []STKey{STKey(rune('s')), STKey(rune('s')), STKey(rune('i')), STKey(rune('s'))}
	result = searcher.find(ssis)
	expectedSSIS := []int64{2}
	if !reflect.DeepEqual(result, expectedSSIS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSSIS)
	}
	ssip := []STKey{STKey(rune('s')), STKey(rune('s')), STKey(rune('i')), STKey(rune('p'))}
	result = searcher.find(ssip)
	expectedSSIP := []int64{5}
	if !reflect.DeepEqual(result, expectedSSIP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSSIP)
	}
	siss := []STKey{STKey(rune('s')), STKey(rune('i')), STKey(rune('s')), STKey(rune('s'))}
	result = searcher.find(siss)
	expectedSISS := []int64{3}
	if !reflect.DeepEqual(result, expectedSISS) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSISS)
	}
	sipp := []STKey{STKey(rune('s')), STKey(rune('i')), STKey(rune('p')), STKey(rune('p'))}
	result = searcher.find(sipp)
	expectedSIPP := []int64{6}
	if !reflect.DeepEqual(result, expectedSIPP) {
		t.Errorf("Find failed, got %s, want %s", result, expectedSIPP)
	}

	ppiDollar := []STKey{STKey(rune('p')), STKey(rune('p')), STKey(rune('i')), STKey(rune('$'))}
	result = searcher.find(ppiDollar)
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

func convertToSTKey(runes []rune) []STKey {
	result := []STKey{}
	for i := 0; i < len(runes); i++ {
		result = append(result, STKey(runes[i]))
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

		dataSource := NewRuneDataSource(s)
		ukkonen := NewUkkonen(dataSource)
		ukkonen.drainDataSource()
		ukkonen.Finish()
		st := ukkonen.Tree()
		//treePrintWithTitle(fmt.Sprintf("Test '%s'", string(s)), st.Root(), ukkonen.Location())
		TreeCheck(st.Root(), dataSource)
		for i := int64(0); i < int64(len(s)); i++ {
			searcher := NewSearcher(st.Root(), dataSource)
			searchFor := s[i:len(s)]
			fmt.Printf("search for suffix %s\n", string(searchFor))
			result := searcher.find(convertToSTKey(searchFor))
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
		searchFor := []STKey{}
		searchFor = append(searchFor, STKey(rune('B')))
		result := searcher.find(searchFor)
		if len(result) == 0 {
			t.Errorf("Expected >1 result, got %d", len(result))
		} else if result[0] != 0 {
			t.Errorf("Expected offset 0, got %d", result[0])
		}*/

	}
}
