package leetcode

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		out    int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, out: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, out: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, out: 4},
		{nums: []int{1, 3, 5, 6}, target: 0, out: 0},
	}

	for _, c := range cases {
		r := searchInsert(c.nums, c.target)
		if r != c.out {
			t.Errorf("in: (%v, %d), expected: %d, actual: %d", c.nums, c.target, c.out, r)
		}
	}
}
