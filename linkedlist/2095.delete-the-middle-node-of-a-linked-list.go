package linkedlist

import (
	"github.com/Ruio9244/leetcode/datastructures/linkedlist"
)

func deleteMiddle(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	p1 := head
	p2 := head.Next
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	p1.Next = p1.Next.Next
	return head
}
