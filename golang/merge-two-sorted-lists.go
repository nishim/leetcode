package leetcode

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var n *ListNode
	if l1.Val <= l2.Val {
		n = &ListNode{l1.Val, nil}
		l1 = l1.Next
	} else {
		n = &ListNode{l2.Val, nil}
		l2 = l2.Next
	}
	root := n

	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			n = appendNode(n, l1.Val)
			l1 = l1.Next
		} else {
			n = appendNode(n, l2.Val)
			l2 = l2.Next
		}
	}

	return root
}

func appendNode(l *ListNode, v int) *ListNode {
	node := &ListNode{v, nil}
	l.Next = node
	return node
}
