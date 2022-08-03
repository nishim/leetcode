package leetcode

import (
	"testing"
)

func TestSqrtx(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{in: 4, out: 2},
		{in: 8, out: 2},
		{in: 9, out: 3},
		{in: 5000, out: 70},
	}

	for _, c := range cases {
		r := mySqrt(c.in)
		if r != c.out {
			t.Errorf("in: %d, expected: %d, actual: %d", c.in, c.out, r)
		}
	}
}
