package bench

import (
	"container/heap"
	"strings"
)

const QueryString = "abbccddeerrffgghhttiiyyoouuppjjllkk"

func ReorganizeString(initialOrder string) string {
	freqMap := make(map[rune]int)

	for _, c := range initialOrder {
		freq := 1

		if v, ok := freqMap[c]; ok {
			freq = v + 1
		}

		if freq > (len(initialOrder)+1)>>1 {
			continue
		}

		freqMap[c] = freq
	}

	h := &MaxHeapQuery{}
	heap.Init(h)

	for k, v := range freqMap {
		heap.Push(h, &Query{R: k, Frequency: v})
	}

	var result strings.Builder

	for !h.Empty() {
		first := heap.Pop(h).(*Query)

		if len(result.String()) == 0 || first.R != rune(result.String()[result.Len()-1]) {
			result.WriteString(string(first.R))
			first.Frequency--

			if first.Frequency > 0 {
				heap.Push(h, first)
			}

		} else {
			if h.Len() == 0 {
				continue
			}

			second := heap.Pop(h).(*Query)
			result.WriteString(string(second.R))
			second.Frequency--

			if second.Frequency > 0 {
				heap.Push(h, second)
			}

			heap.Push(h, first)
		}
	}

	if len(initialOrder) != result.Len() {
		return ""
	}

	return result.String()
}

type Query struct {
	R         rune
	Frequency int
}

type MaxHeapQuery []*Query

func (h *MaxHeapQuery) Len() int {
	return len(*h)
}

func (h *MaxHeapQuery) Less(i, j int) bool {
	return (*h)[i].Frequency > (*h)[j].Frequency
}

func (h *MaxHeapQuery) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeapQuery) Push(x interface{}) {
	*h = append(*h, x.(*Query))
}

func (h *MaxHeapQuery) Top() *Query {
	return (*h)[0]
}

func (h *MaxHeapQuery) Empty() bool {
	return len(*h) == 0
}

func (h *MaxHeapQuery) Pop() interface{} {
	var val *Query

	n := len(*h)
	val, *h = (*h)[n-1], (*h)[0:n-1]

	return val
}
