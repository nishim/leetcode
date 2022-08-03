package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		in  []string
		out string
	}{
		{in: []string{"flower", "flow", "flight"}, out: "fl"},
		{in: []string{"dog", "racecar", "car"}, out: ""},
	}

	for _, c := range cases {
		r := longestCommonPrefix(c.in)
		if r != c.out {
			t.Errorf("expected: %s, actual: %s", c.out, r)
		}
	}
}
