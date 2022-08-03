package leetcode

import (
	"testing"
)

func TestRpeatedStringMatch(t *testing.T) {
	cases := []struct {
		a   string
		b   string
		out int
	}{
		{a: "abcd", b: "cdabcdab", out: 3},
		{a: "a", b: "aa", out: 2},
		{a: "a", b: "a", out: 1},
		{a: "abc", b: "wxyz", out: -1},
		{a: "aaaaaaaaaaaaaaaaaaaaaab", b: "ba", out: 2},
	}

	for _, c := range cases {
		r := repeatedStringMatch(c.a, c.b)
		if r != c.out {
			t.Errorf("in: %s, %s, expected: %d, actual: %d", c.a, c.b, c.out, r)
		}
	}
}
