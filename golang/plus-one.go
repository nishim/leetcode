package leetcode

func plusOne(digits []int) []int {
	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + c
		if digits[i] == 10 {
			c = 1
			digits[i] = 0
		} else {
			c = 0
			break
		}
	}

	if c == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
