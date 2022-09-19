package bench

import (
	"container/heap"
	"strconv"
	"strings"
)

func ReorganizeStringSlice(initialOrder string) string {
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

	h := &MaxHeap{}
	heap.Init(h)

	for k, v := range freqMap {
		heap.Push(h, []int{int(k), v})
	}

	var result strings.Builder

	for !h.Empty() {
		first := heap.Pop(h).([]int)

		if result.Len() == 0 || first[0] != int(result.String()[result.Len()-1]) {
			result.WriteString(strconv.Itoa(first[0]))
			first[1]--

			if first[1] > 0 {
				heap.Push(h, first)
			}
		} else {
			if h.Len() == 0 {
				continue
			}

			second := heap.Pop(h).([]int)

			result.WriteString(strconv.Itoa(second[0]))
			second[1]--

			if second[1] > 0 {
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

type MaxHeap [][]int

func (h *MaxHeap) Len() int { return len(*h) }

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i][1] > (*h)[j][1]
}

func (h *MaxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Top() []int {
	return (*h)[0]
}

func (h *MaxHeap) Empty() bool {
	return len(*h) == 0
}

func (h *MaxHeap) Pop() interface{} {
	var val []int

	n := len(*h)
	val, *h = (*h)[n-1], (*h)[0:n-1]
	return val
}
