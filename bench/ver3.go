package bench

import "container/heap"

func reorganizeStringVer3(s string) string {
	countMap := make(map[string]int)

	for _, ch := range s {
		countMap[string(ch)] += 1
	}

	h := &maxHeap{}

	for v, c := range countMap {
		heap.Push(h, char{v, c})
	}

	var res string

	for h.Len() > 0 {
		firstChar := heap.Pop(h).(char)

		if h.Len() == 0 && firstChar.count > 1 {
			return ""
		} else if h.Len() == 0 && firstChar.count == 1 {
			return res + firstChar.value
		}

		secondChar := heap.Pop(h).(char)
		res = res + firstChar.value + secondChar.value

		if firstChar.count-1 > 0 {
			heap.Push(h, char{firstChar.value, firstChar.count - 1})
		}

		if secondChar.count-1 > 0 {
			heap.Push(h, char{secondChar.value, secondChar.count - 1})
		}
	}

	return res
}

type char struct {
	value string
	count int
}

type maxHeap []char

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	val := x.(char)
	*h = append(*h, val)
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
