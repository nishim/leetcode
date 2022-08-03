package leetcode

import (
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	if b == "" {
		return 0
	}

	i := 1
	for s := a; len(s) <= (2*len(b)) || i <= 2; s = s + a {
		if strings.Index(s, b) >= 0 {
			return i
		}
		i++
	}

	return -1
}
