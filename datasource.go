package suffixtree

import (
	"fmt"
	"io"
	"os"
	"log"
	"bufio"
)

// A DataSource provides a sequence of STKey values over a channel, and allows individual STKey values
// to be retrieved by their offset.
type DataSource interface {
	KeyAtOffset(int64) STKey
	STKeys() <-chan STKey
	StringFrom(start, end int64) string
}

type stringDataSource struct {
	runes  []rune
	stream <-chan STKey
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

func (dataSource *stringDataSource) KeyAtOffset(offset int64) STKey {
	return STKey(dataSource.runes[offset])
}

func (dataSource *stringDataSource) STKeys() <-chan STKey {
	return dataSource.stream
}

func (s *stringDataSource) StringFrom(start, end int64) string {
	x := ""
	if end < 0 {
		end = start
		x = "..."
	}
	result := ""
	for start <= end {
		result = fmt.Sprintf("%s%c", result, s.KeyAtOffset(start))
		start++
	}
	result = fmt.Sprintf("%s%s", result, x)
	return result
}

type fileDataSource struct {
	positionalReader *os.File
	stream <-chan STKey
	singleByte []byte
}

func NewFileDataSource(filePath string) DataSource {
	dataChannel := make(chan STKey, 1024)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("File not available")
	}
	sequentialReader := bufio.NewReader(file)
	positionalReader, err := os.Open(filePath)
	go func(reader io.Reader, dataChannel chan<- STKey) {
		var singleByte []byte = []byte{0}
		var err error
		for {
			_, err = reader.Read(singleByte)
			if err == io.EOF {
				break
			}
			dataChannel <- STKey(singleByte[0])
		}
		close(dataChannel)
	}(sequentialReader, dataChannel)
	return &fileDataSource{positionalReader, dataChannel, []byte{0}}
}

func (f *fileDataSource) KeyAtOffset(offset int64) STKey {
	f.positionalReader.Seek(offset, os.SEEK_SET)
	f.positionalReader.Read(f.singleByte)
	return STKey(f.singleByte[0])
}

func (f *fileDataSource) STKeys() <-chan STKey {
	return f.stream
}

func (f *fileDataSource) StringFrom(start, end int64) string {
	var byteArray = make([]byte, end - start + 1)
	f.positionalReader.Read(byteArray)
	return string(byteArray)
}