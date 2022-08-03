package leetcode

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		digits []int
		out    []int
	}{
		{digits: []int{1, 2, 3}, out: []int{1, 2, 4}},
		{digits: []int{4, 3, 2, 1}, out: []int{4, 3, 2, 2}},
		{digits: []int{0}, out: []int{1}},
		{digits: []int{9, 9}, out: []int{1, 0, 0}},
	}

	for _, c := range cases {
		r := plusOne(c.digits)
		for i, v := range r {
			if v != c.out[i] {
				t.Errorf("expected: %d, actual: %d", c.out, r)
			}
		}
	}
}
