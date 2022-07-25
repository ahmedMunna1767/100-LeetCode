package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	1 -> 2 -> 3 -> 4 -> 5

	Step: 01
	current: 1 -> -> 5
	prev: nil

	Step: 02
	current: 2 -> -> 5
	prev: 1 --> (current.next = previous)

	Step: 03
	current: 3 -> -> 5
	prev: 2 -> 1 --> (current.next = previous)

	Step: 04
	current: 4 -> -> 5
	prev: 3 -> 2 -> 1 (current.next = previous)

	Step: 05
	current: 5
	prev: 4 -> 3 -> 2 -> 1 (current.next = previous)

	Step: 06
	current: current = head
			 current.next = prev

*/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var reverseUtil func(current, previous *ListNode)

	reverseUtil = func(current, previous *ListNode) {
		if current.Next == nil {
			head = current
			current.Next = previous
			return
		}
		next := current.Next
		current.Next = previous
		reverseUtil(next, current)
	}

	reverseUtil(head, nil)

	return head
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}

func ReverseList() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}

	ret := reverseList(&head)

	printLinkedList(ret)

}
