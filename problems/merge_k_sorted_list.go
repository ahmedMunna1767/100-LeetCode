package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKSortedLists(lists []*ListNode) *ListNode {

	var rootNode ListNode
	if len(lists) == 0 {
		return nil
	}

	shouldBreak := true
	for _, node := range lists {
		if node != nil {
			shouldBreak = false
			break
		}
	}
	if shouldBreak {
		return nil
	}
	var curNode *ListNode
	first := true
	for {
		min := 100000000
		curMinIndex := 10000000
		shouldBreak := true

		for i := range lists {
			if lists[i] == nil {
				continue
			}
			if min > lists[i].Val {
				min = lists[i].Val
				curMinIndex = i
			}
		}

		if first {
			rootNode = ListNode{
				Val: min,
			}
			first = false
			curNode = &rootNode
		} else {
			curNode.Next = &ListNode{
				Val: min,
			}
			curNode = curNode.Next
		}

		if lists[curMinIndex].Next != nil {
			lists[curMinIndex] = lists[curMinIndex].Next
		} else {
			lists[curMinIndex] = nil
		}

		for _, node := range lists {
			if node != nil {
				shouldBreak = false
			}
		}

		if shouldBreak {
			break
		}

	}
	return &rootNode
}
