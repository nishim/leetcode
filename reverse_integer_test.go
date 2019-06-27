package leetcode

import (
	"testing"
)

func TestReverse(t *testing.T) {
	r := reverse(123)
	if r != 321 {
		t.Error("Error")
	}

	r = reverse(1534236469)
	if r != 0 {
		t.Error("Error")
	}
}
