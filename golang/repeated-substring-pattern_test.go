package leetcode

import (
	"testing"
)

func TestRpeatedSubstringPattern(t *testing.T) {
	cases := []struct {
		s   string
		out bool
	}{
		{s: "bb", out: true},
		{s: "abab", out: true},
		{s: "aba", out: false},
		{s: "abcabcabcabc", out: true},
	}

	for _, c := range cases {
		r := repeatedSubstringPattern(c.s)
		if r != c.out {
			t.Errorf("in: %s, expected: %t, actual: %t", c.s, c.out, r)
		}
	}
}
