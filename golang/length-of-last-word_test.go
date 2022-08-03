package leetcode

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	cases := []struct {
		s   string
		out int
	}{
		{s: "hello world", out: 5},
		{s: "good evening", out: 7},
		{s: "good bye", out: 3},
		{s: "a ", out: 1},
		{s: "b   a    ", out: 1},
		{s: "", out: 0},
		{s: " ", out: 0},
	}

	for _, c := range cases {
		r := lengthOfLastWord(c.s)
		if r != c.out {
			t.Errorf("in: %s, expected: %d, actual: %d", c.s, c.out, r)
		}
	}
}
