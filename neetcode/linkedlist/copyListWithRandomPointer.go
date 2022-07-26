package linkedlist

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)

	if head == nil {
		return nil
	}
	copyOrg := head
	copyHead := &Node{Val: head.Val, Next: nil}
	curNode := copyHead
	nodeMap[head] = copyHead

	for copyOrg.Next != nil {
		curNode.Next = &Node{Val: copyOrg.Next.Val, Next: nil}
		nodeMap[copyOrg.Next] = curNode.Next
		curNode = curNode.Next
		copyOrg = copyOrg.Next
	}
	curNode = copyHead
	copyOrg = head

	for copyOrg != nil {
		if copyOrg.Random == nil {
			curNode.Random = nil
		} else {
			curNode.Random = nodeMap[copyOrg.Random]
		}

		copyOrg = copyOrg.Next
		curNode = curNode.Next
	}
	return copyHead
}

func CopyRandomList() {
	head := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
				Next: &Node{
					Val:  4,
					Next: &Node{Val: 5},
				},
			},
		},
	}

	copy := copyRandomList(head)

	for copy != nil {
		print(copy.Val)
		copy = copy.Next
	}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
