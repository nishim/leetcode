package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	crr := 0

	for l1 != nil || l2 != nil || crr > 0 {
		v := crr
		crr = 0

		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		if v >= 10 {
			crr = 1
			v = v - 10
		}

		cur.Next = &ListNode{v, nil}
		cur = cur.Next
	}

	return root.Next
}
