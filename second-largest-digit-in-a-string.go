package leetcode

func secondHighest(s string) int {
	h1 := -1
	h2 := -1
	for _, r := range s {
		if '0' > r || '9' < r {
			continue
		}

		i := int(r - '0')

		if i > h1 {
			h2 = h1
			h1 = i
		} else if i > h2 && i != h1 {
			h2 = i
		}

		if h2 == 8 {
			return 8
		}
	}
	return h2
}
