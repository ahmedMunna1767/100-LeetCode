package linkedlist

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var prevTarget, headCopy *ListNode
	target := head
	headCopy = head
	count := 0

	for headCopy != nil {
		fmt.Println(target.Val)
		count = count + 1
		if count > n {
			prevTarget = target
			target = target.Next
		}
		headCopy = headCopy.Next
	}

	if target == head {
		return head.Next
	} else {
		prevTarget.Next = target.Next
		return head
	}
}

func RemoveNthFromEnd() {
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

	printLinkedList(removeNthFromEnd(&head, 4))
}
