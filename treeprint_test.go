package suffixtree

import "testing"

func TestStringTree(t *testing.T) {
	testString := "mississippi"
	runes := []rune(testString)
	s := suffixtree.NewStringDataSource(testString)
	incomingChannel := s.STKeys()
	for _, r := range runes {
		test := <-incomingChannel
		if test != suffixtree.STKEY(r) {
			t.Error("channel did not provide the expected value")
		}
	}
}
