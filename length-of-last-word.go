package leetcode

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	spl := strings.Split(strings.TrimSpace(s), " ")
	return len(spl[len(spl)-1])
}
