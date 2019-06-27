package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	r := twoSum([]int{2, 7, 11, 15}, 9)
	if !(r[0] == 0 && r[1] == 1) {
		t.Error("Error")
	}
}
