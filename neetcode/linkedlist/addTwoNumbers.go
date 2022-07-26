package linkedlist

import (
	"math"
)

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := 0, 0
	pos := 0

	for l1 != nil {
		num1 = num1 + l1.Val*(int(math.Pow10(pos)))
		l1 = l1.Next
		pos++
	}

	pos = 0
	for l2 != nil {
		num2 = num2 + l2.Val*(int(math.Pow10(pos)))
		l2 = l2.Next
		pos++
	}

	num1 = num1 + num2
	var head, current *ListNode
	if num1 == 0 {
		return &ListNode{Val: 0}
	}

	for num1 > 0 {
		digit := num1 % 10
		num1 = num1 / 10

		if head == nil {
			head = &ListNode{Val: digit}
			current = head
		} else {
			current.Next = &ListNode{Val: digit}
			current = current.Next
		}
	}

	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, current *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		if head == nil {
			head = &ListNode{Val: (l1.Val + l2.Val + carry) % 10}
			current = head
			carry = (l1.Val + l2.Val + carry) / 10
		} else {
			current.Next = &ListNode{Val: (l1.Val + l2.Val + carry) % 10}
			current = current.Next
			carry = (l1.Val + l2.Val + carry) / 10
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		current.Next = &ListNode{Val: (l1.Val + carry) % 10}
		current = current.Next
		carry = (l1.Val + carry) / 10
		l1 = l1.Next
	}

	for l2 != nil {
		current.Next = &ListNode{Val: (l2.Val + carry) % 10}
		current = current.Next
		carry = (l2.Val + carry) / 10
		l2 = l2.Next
	}

	if carry != 0 {
		current.Next = &ListNode{Val: carry}
	}
	return head
}

func AddTwoNumbers() {
	list1 := ListNode{
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

	list2 := ListNode{
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

	printLinkedList(addTwoNumbers2(&list1, &list2))
	addTwoNumbers(&list1, &list2)
}
