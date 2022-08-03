package leetcode

import "fmt"

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	max := ""

	for i := 0; i < l; i++ {
		used := make(map[string]struct{})
		for j := i + 1; j <= l; j++ {
			fmt.Println(i, j)
			c := s[(j - 1):j]
			if _, ok := used[c]; ok {
				sub := s[i:(j - 1)]
				if len(max) < len(sub) {
					max = sub
				}
				break
			}

			used[c] = struct{}{}

			if j == l && len(max) < len(s[i:j]) {
				max = s[i:j]
			}
		}
	}

	return len(max)
}
