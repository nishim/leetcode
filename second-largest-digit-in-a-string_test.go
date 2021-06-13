package leetcode

import (
	"testing"
)

func TestSecondLargestDigitInAString(t *testing.T) {
	cases := []struct {
		s   string
		out int
	}{
		{s: "dfa12321afd", out: 2},
		{s: "abc1111", out: -1},
	}

	for _, c := range cases {
		r := secondHighest(c.s)
		if r != c.out {
			t.Errorf("in: %s, expected: %d, actual: %d", c.s, c.out, r)
		}
	}
}
