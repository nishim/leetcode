package leetcode

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		nums []int
		val  int
		out1 int
		out2 []int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, out1: 2, out2: []int{2, 2}},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, out1: 5, out2: []int{0, 1, 4, 0, 3}},
	}

	for _, c := range cases {
		nums := c.nums
		r := removeElement(nums, c.val)
		if r != c.out1 {
			t.Errorf("expected: (%d, %v), actual: (%d, %v)", c.out1, c.out2, r, nums)
		}
	}
}
