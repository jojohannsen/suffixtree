package suffixtree

import "fmt"

type BoundaryMap interface {
	Append(segment Segment)
	IncludesAllSegments(values []int32) bool
	Name() string
	Segments() []Segment
}

// an arbitrary value belongs to a specific Segment if it falls between the min and max boundary (inclusive)
type Segment struct {
	minBoundary int64
	maxBoundary int64
	name        string
}

func (s *Segment) Includes(val int32) bool {
	v64 := int64(val)
	return v64 >= s.minBoundary && v64 <= s.maxBoundary
}

func (s *Segment) String() string {
	return fmt.Sprintf("%s,%d,%d\n", s.name, s.minBoundary, s.maxBoundary)
}

type boundaryMap struct {
	name string
	segments []Segment
}

func (bm *boundaryMap) Segments() []Segment {
	return bm.segments
}

func (bm *boundaryMap) Name() string {
	return bm.name
}

func (bm *boundaryMap) Append(segment Segment) {
	bm.segments = append(bm.segments, segment)
}

func (bm *boundaryMap) Dump(boundaryName string) {
	fmt.Printf("Boundary: %s, %d sections\n", boundaryName, len(bm.segments))
	for _, segment := range bm.segments {
		fmt.Printf("  %s bounds: %d to %d\n", segment.name, segment.minBoundary, segment.maxBoundary)
	}
	fmt.Printf("End of boundary %s\n", boundaryName)
}

func (bm *boundaryMap) IncludesAllSegments(values []int32) bool {
	included := make([]bool, len(bm.segments))
	for _, val := range values {
		for i, segment := range bm.segments {
			if !included[i] {
				included[i] = segment.Includes(val)
			}
		}
	}
	for _, includesVal := range included {
		if !includesVal {
			return false
		}
	}
	return true
}

func NewBoundaryMap(boundaryName string, suffixTree SuffixTree) BoundaryMap {
	searcher := NewSearcher(suffixTree.Root(), suffixTree.DataSource())
	findStr := "$" + boundaryName + "$"
	findSTKey := []STKey{}
	for _, c := range findStr {
		findSTKey = append(findSTKey, STKey(c))
	}
	result := searcher.Find(findSTKey)
	boundaryMapResult := &boundaryMap{name: boundaryName}
	nextStartBoundary := int64(1)
	for _, offset := range result {
		startBoundary := nextStartBoundary - 1
		endBoundary := int64(offset) - 1
		name := suffixTree.DataSource().StringFromTo(offset+int32(len(findStr)), "$")
		boundaryMapResult.Append(Segment{startBoundary, endBoundary, name})

		nextStartBoundary = endBoundary + int64(len(findStr)+len(name)+3)

		// if this boundary is actually a higher level boundary, bypass that
		testName := suffixTree.DataSource().StringFrom(int32(nextStartBoundary), int32(nextStartBoundary + 8))
		fmt.Printf("TestName='%s' from %d\n", testName, nextStartBoundary)
		if testName == "$species$" {
			nextStartBoundary += 8
			speciesName := suffixTree.DataSource().StringFromTo(int32(nextStartBoundary + 1), "$")
			fmt.Printf("nextStartBoundary now %d, speciesName='%s'\n", nextStartBoundary, speciesName)
			nextStartBoundary += int64(len(speciesName) + 3)
		}
		//fmt.Printf("boundary %s at %d is %s\n", boundaryName, offset, suffixTree.DataSource().StringFrom(offset, offset+30))
		//fmt.Printf("%d=%s\n", offset + int32(len(findStr)), name)
	}

	boundaryMapResult.Dump(boundaryName)
	return boundaryMapResult
}
