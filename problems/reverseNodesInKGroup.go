package problems

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var current, next, prev, nl ListNode
	current = *head
	count := 0

	for current != nl && count < k {
		if current.Next != nil {
			next = *current.Next
		}
		current.Next = &prev
		prev = current
		current = next
		count++
		fmt.Println(count)
	}

	if next != nl {
		head.Next = ReverseKGroup(&next, k)
	}
	return &prev
}

func TestReverseGroup() {
	listOne := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	node := ReverseKGroup(&listOne, 3)
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
}
