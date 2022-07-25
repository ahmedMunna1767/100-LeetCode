package linkedlist

func reorderList(head *ListNode) *ListNode {

	var findMiddle func(tortoise, hare *ListNode) *ListNode
	var reverse func(head *ListNode) *ListNode

	// Tortoise & Hare method to find middle of a linked list
	findMiddle = func(tortoise, hare *ListNode) *ListNode {
		for hare != nil && hare.Next != nil {
			tortoise = tortoise.Next
			hare = hare.Next.Next
		}
		return tortoise
	}

	// reverse a linked list
	reverse = func(head *ListNode) *ListNode {
		var prev, next *ListNode
		current := head
		for current != nil {
			next = current.Next
			current.Next = prev
			prev = current
			current = next
		}
		return prev
	}

	middle := findMiddle(head, head.Next)
	firstHalf := head

	lastHalf := reverse(middle.Next)
	middle.Next = nil
	// list is divided into two half lists

	for firstHalf != nil && lastHalf != nil {
		last := lastHalf
		lastHalf = lastHalf.Next
		last.Next = firstHalf.Next
		firstHalf.Next = last
		firstHalf = last.Next
	}
	return head
}

func ReorderList() {
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}},
				},
			},
		},
	}

	printLinkedList(reorderList(&list1))
}
