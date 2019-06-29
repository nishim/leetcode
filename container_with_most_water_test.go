package leetcode

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	cases := []struct {
		in  []int
		out int
	}{
		{in: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, out: 49},
	}

	for _, c := range cases {
		r := maxArea(c.in)
		if r != c.out {
			t.Errorf("expected: %d, actual: %d", c.out, r)
		}
	}
}
