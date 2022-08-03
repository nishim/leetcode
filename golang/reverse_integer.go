package leetcode

// https://leetcode.com/problems/reverse-integer/

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x = x * -1
	}

	r := ""
	for _, v := range strconv.Itoa(x) {
		r = string(v) + r
	}

	i, _ := strconv.Atoi(r)
	if negative {
		i = i * -1
	}
	if int(math.Pow(2, 31)-1) < i || int(math.Pow(-2, 31)) > i {
		return 0
	}
	return i
}
