package linkedlist

import (
	"github.com/Ruio9244/leetcode/datastructures/linkedlist"
)

func middleNode(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := head
	length := 1
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		length = length + 2
	}
	if p2.Next != nil {
		length++
	}
	if length%2 == 0 {
		return p1.Next
	}
	return p1
}
