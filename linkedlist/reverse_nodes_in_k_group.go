package linkedlist

// 25. Reverse Nodes in k-Group
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	cur, n := head, 0
	for cur != nil {
		n++
		cur = cur.Next
	}

	dummy := &ListNode{}
	tail := dummy

	for i := 0; i < n/k; i++ {
		nextTailPosition := head
		for j := 0; j < k; j++ {
			next := head.Next
			head.Next = tail.Next
			tail.Next = head
			head = next
		}
		tail = nextTailPosition
	}
	// 链接剩下的节点
	tail.Next = head
	return dummy.Next
}
