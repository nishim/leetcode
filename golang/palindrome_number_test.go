package leetcode

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{in: 121, out: true},
		{in: -121, out: false},
		{in: 10, out: false},
	}

	for _, c := range cases {
		r := isPalindrome(c.in)
		if r != c.out {
			t.Errorf("expected: %t, actual: %t", c.out, r)
		}
	}
}
