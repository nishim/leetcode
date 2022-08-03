package leetcode

func intToRoman(num int) string {
	roman := ""
	a := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	b := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for num > 0 {
		for pos, i := range a {
			if num >= i {
				num = num - i
				roman = roman + b[pos]
				break
			}
		}
	}
	return roman
}
