package priorityQueue

import (
	"fmt"

	"github.com/ahmedMunna1767/100-LeetCode/myHeap"
)

type Item struct {
	Value    string
	Index    int
	Priority int
}

func (item Item) Print() {
	fmt.Printf("(%d). %.2d:%s ", item.Index, item.Priority, item.Value)
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// place elements with higher priority at the front
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	myHeap.Fix(pq, item.Index)
}

func (pq PriorityQueue) Print() {
	for _, item := range pq {
		item.Print()
	}
	println()
}

func TestMyPriorityQueue() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}

	myHeap.Init(&pq)
	pq.Print()
	item := &Item{
		Value:    "orange",
		Priority: 1,
	}

	myHeap.Push(&pq, item)
	pq.Print()
	pq.Update(item, item.Value, 5)
	pq.Print()
	for pq.Len() > 0 {
		item := myHeap.Pop(&pq).(*Item)
		item.Print()
		println()
	}
}
