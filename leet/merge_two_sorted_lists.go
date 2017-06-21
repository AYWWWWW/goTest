package leet

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{0, nil}
	next := &head
	head1, head2 := l1, l2
	for head1 != nil && head2 != nil {
		if head1.Val > head2.Val {
			next.Next = head2
			head2 = head2.Next
		} else {
			next.Next = head1
			head1 = head1.Next
		}
		next = next.Next
	}
	if head1 != nil {
		next.Next = head1
	}
	if head2 != nil {
		next.Next = head2
	}
	return head.Next
}
