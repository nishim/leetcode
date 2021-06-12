package leetcode

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	hlen := len(haystack)
	nlen := len(needle)

	if nlen > hlen {
		return -1
	}

	for i := 0; (i + nlen - 1) < hlen; i++ {
		if haystack[i] != needle[0] {
			continue
		}

		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}
