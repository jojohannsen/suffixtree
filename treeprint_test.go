package suffixtree

import "testing"

func TestStringTree(t *testing.T) {
	testString := "mississippi"
	runes := []rune(testString)
	s := NewStringDataSource(testString)
	incomingChannel := s.STKeys()
	for _, r := range runes {
		test := <-incomingChannel
		if test != STKey(r) {
			t.Error("channel did not provide the expected value")
		}
	}
}
