package leetcode

import (
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	cases := []struct {
		in1 *ListNode
		in2 *ListNode
		out *ListNode
	}{
		{
			in1: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			in2: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			out: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		}, {
			in1: nil, in2: nil, out: nil,
		}, {
			in1: nil,
			in2: &ListNode{0, nil},
			out: &ListNode{0, nil},
		},
	}

	for _, c := range cases {
		r := mergeTwoLists(c.in1, c.in2)
		a := r
		e := c.out
		for {
			if a == nil && e == nil {
				break
			}
			if a == nil || e == nil {
				t.Errorf("in: %v, %v, expected: %v, actual: %v", c.in1, c.in2, c.out, r)
				break
			}
			if a.Val != e.Val {
				t.Errorf("in: %v, %v, expected: %v, actual: %v", c.in1, c.in2, c.out, r)
				break
			}
			a = a.Next
			e = e.Next
		}
	}
}
