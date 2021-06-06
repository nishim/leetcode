package leetcode

import (
	"testing"
)

func TestIntegerToRoman(t *testing.T) {
	cases := []struct {
		num int
		out string
	}{
		{num: 3, out: "III"},
		{num: 4, out: "IV"},
		{num: 9, out: "IX"},
		{num: 20, out: "XX"},
		{num: 58, out: "LVIII"},
		{num: 1994, out: "MCMXCIV"},
	}

	for _, c := range cases {
		r := intToRoman(c.num)
		if r != c.out {
			t.Errorf("in: %d, expected: %s, actual: %s", c.num, c.out, r)
		}
	}
}
