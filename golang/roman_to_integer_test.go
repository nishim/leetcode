package leetcode

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{in: "III", out: 3},
		{in: "IV", out: 4},
		{in: "IX", out: 9},
		{in: "LVIII", out: 58},
		{in: "MCMXCIV", out: 1994},
		{in: "MDCCCLXXXIV", out: 1884},
	}

	for _, c := range cases {
		r := romanToInt(c.in)
		if r != c.out {
			t.Errorf("expected: %d, actual: %d", c.out, r)
		}
	}
}
