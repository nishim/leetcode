package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		l1  *ListNode
		l2  *ListNode
		out *ListNode
	}{
		{
			l1:  &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			l2:  &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			out: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		}, {
			l1:  &ListNode{0, nil},
			l2:  &ListNode{0, nil},
			out: &ListNode{0, nil},
		}, {
			l1:  &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			l2:  &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			out: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
	}

	for _, c := range cases {
		r := addTwoNumbers(c.l1, c.l2)
		a := r
		e := c.out
		for {
			if a == nil && e == nil {
				break
			}
			if a == nil || e == nil {
				t.Errorf("l: %v, %v, expected: %v, actual: %v", c.l1, c.l2, c.out, r)
				break
			}
			if a.Val != e.Val {
				t.Errorf("l: %v, %v, expected: %v, actual: %v", c.l1, c.l2, c.out, r)
				break
			}
			a = a.Next
			e = e.Next
		}
	}
}
