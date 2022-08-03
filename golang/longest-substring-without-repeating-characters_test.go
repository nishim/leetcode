package leetcode

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	cases := []struct {
		s   string
		out int
	}{
		{s: "abcabcbb", out: 3},
		{s: "bbbbb", out: 1},
		{s: "pwwkew", out: 3},
		{s: "", out: 0},
		{s: " ", out: 1},
		{s: "au", out: 2},
	}

	for _, c := range cases {
		r := lengthOfLongestSubstring(c.s)
		if r != c.out {
			t.Errorf("in: %s, expected: %d, actual: %d", c.s, c.out, r)
		}
	}
}
