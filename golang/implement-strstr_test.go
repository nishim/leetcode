package leetcode

import (
	"testing"
)

func TestImplementStrStr(t *testing.T) {
	cases := []struct {
		heystack string
		needle   string
		out      int
	}{
		{heystack: "hello", needle: "ll", out: 2},
		{heystack: "aaaaa", needle: "bba", out: -1},
		{heystack: "a", needle: "a", out: 0},
		{heystack: "", needle: "", out: 0},
	}

	for _, c := range cases {
		r := strStr(c.heystack, c.needle)
		if r != c.out {
			t.Errorf("in: %s, %s, expected: %d, actual: %d", c.heystack, c.needle, c.out, r)
		}
	}
}
