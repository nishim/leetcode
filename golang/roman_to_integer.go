package leetcode

// https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
	num := 0
	before := ""
	for i := len(s) - 1; i >= 0; i-- {
		switch string(s[i]) {
		case "I":
			if before == "V" || before == "X" {
				num = num - 1
			} else {
				num = num + 1
			}
		case "V":
			num = num + 5
		case "X":
			if before == "L" || before == "C" {
				num = num - 10
			} else {
				num = num + 10
			}
		case "L":
			num = num + 50
		case "C":
			if before == "D" || before == "M" {
				num = num - 100
			} else {
				num = num + 100
			}
		case "D":
			num = num + 500
		case "M":
			num = num + 1000

		}

		before = string(s[i])
	}

	return num
}
