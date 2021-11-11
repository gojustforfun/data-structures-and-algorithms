package linkedlist

// 21. Merge Two Sorted Lists
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 
 type ListNode struct {
	Val int
	Next *ListNode
 }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for l1 != nil || l2 != nil {
		// 要L1上的元素的条件:
		// 1) L2为空 l2 == nil
		// 2) L1和L2均不为空且L1.Val小于等于L2.Val
		// 此题与 ../array/merge_sorted_array.go的解题思路类似 
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	return head.Next
}