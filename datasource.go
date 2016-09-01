package suffixtree

import "fmt"

// A DataSource provides a sequence of STKey values over a channel, and allows individual STKey values
// to be retrieved by their offset.
type DataSource interface {
	keyAtOffset(int64) STKey
	STKeys() chan STKey
	stringFrom(start, end int64) string
}

func (s *stringDataSource) stringFrom(start, end int64) string {
	x := ""
	if end < 0 {
		end = start
		x = "..."
	}
	result := ""
	for start <= end {
		result = fmt.Sprintf("%s%c", result, s.keyAtOffset(start))
		start++
	}
	result = fmt.Sprintf("%s%s", result, x)
	return result
}

type stringDataSource struct {
	runes  []rune
	stream chan STKey
}

func NewRuneDataSource(runes []rune) DataSource {
	dataChannel := make(chan STKey)
	go func(runes []rune, dataChannel chan<- STKey) {
		for _, r := range runes {
			dataChannel <- STKey(r)
		}
		close(dataChannel)
	}(runes, dataChannel)
	return &stringDataSource{runes, dataChannel}
}

func NewStringDataSource(s string) DataSource {
	runes := []rune(s)
	return NewRuneDataSource(runes)
}

func (dataSource *stringDataSource) keyAtOffset(offset int64) STKey {
	return STKey(dataSource.runes[offset])
}

func (dataSource *stringDataSource) STKeys() chan STKey {
	return dataSource.stream
}
