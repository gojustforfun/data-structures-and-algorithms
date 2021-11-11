package linkedlist

// 206. Reverse Linked List
// https://leetcode-cn.com/problems/reverse-linked-list/

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev, cur := (*ListNode)(nil), head

	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

    return prev
}