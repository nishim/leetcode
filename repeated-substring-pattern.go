package leetcode

func repeatedSubstringPattern(s string) bool {
	l := len(s)

	for i := 1; i <= l/2; i++ {
		if l%i != 0 {
			continue
		}

		sub := s[:i]

		repeated := true
		for j := i; j < l; j = j + i {
			if sub != s[j:j+i] {
				repeated = false
				break
			}
		}

		if repeated {
			return true
		}
	}

	return false
}
