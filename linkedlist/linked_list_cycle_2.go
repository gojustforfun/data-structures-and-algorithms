package linkedlist

// 142. Linked List Cycle II
// https://leetcode-cn.com/problems/linked-list-cycle-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	node := cycleNode(head)
	if node == nil {
		return node
	}

	for cur := head; cur != node; {
		cur = cur.Next
		node = node.Next
	}

	return node
}

func cycleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}
