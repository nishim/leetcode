package leetcode

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	cases := []struct {
		in  string
		out bool
	}{
		{in: "()", out: true},
		{in: "()[]{}", out: true},
		{in: "(]", out: false},
		{in: "([)]", out: false},
		{in: "{[]}", out: true},
	}

	for _, c := range cases {
		r := isValid(c.in)
		if r != c.out {
			t.Errorf("in: %s, expected: %t, actual: %t", c.in, c.out, r)
		}
	}
}
