package leetcode

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	r := ""

	for i := 0; ; i++ {
		s := ""
		for _, str := range strs {
			if (len(str) - 1) < i {
				s = ""
				break
			}

			if s == "" {
				s = string(str[i])
			} else if s != string(str[i]) {
				s = ""
				break
			}
		}

		if s == "" {
			break
		}

		r = r + s
	}

	return r
}
