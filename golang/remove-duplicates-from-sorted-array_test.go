package leetcode

import (
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	cases := []struct {
		in  []int
		out int
	}{
		{in: []int{1, 1, 2}, out: 2},
		{in: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, out: 5},
	}

	for _, c := range cases {
		r := removeDuplicates(c.in)
		if r != c.out {
			t.Errorf("in: %v, expected: %v, actual: %v", c.in, c.out, r)
		}
	}
}
