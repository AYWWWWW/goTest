package leet

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	p2 := head
	for n > 0 {
		p1 = p1.Next
		n--
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	if p2 == head {
		return head.Next
	}
	deleteNode(head, p2)
	return head
}

func deleteNode(head, p *ListNode) {
	if p.Next == nil {
		pre := head
		for pre.Next != nil && pre.Next != p {
			pre = pre.Next
		}
		pre.Next = p.Next
	} else {
		p.Val = p.Next.Val
		p.Next = p.Next.Next
	}
}
