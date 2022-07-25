package linkedlist

import "fmt"

func hasCycle(head *ListNode) bool {
	hare, tortoise := head, head.Next

	for hare != nil && hare.Next != nil {
		if hare == tortoise {
			return false
		}
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}
	return true
}

func HasCycle() {
	fmt.Println(hasCycle(&ListNode{}))
}
