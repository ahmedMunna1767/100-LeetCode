package linkedlist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, cur *ListNode

	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list2
	}

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			if head == nil {
				head = &ListNode{Val: list2.Val}
				cur = head
			} else {
				cur.Next = &ListNode{Val: list2.Val}
				cur = cur.Next
			}
			list2 = list2.Next
		} else {
			if head == nil {
				head = &ListNode{Val: list1.Val}
				cur = head
			} else {
				cur.Next = &ListNode{Val: list1.Val}
				cur = cur.Next
			}
			list1 = list1.Next
		}
	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return head
}

func MergeTwoLists() {
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

	printLinkedList(mergeTwoLists(&list1, &list2))

}
