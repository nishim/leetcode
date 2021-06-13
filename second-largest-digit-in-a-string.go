package leetcode

func secondHighest(s string) int {
	h1 := -1
	h2 := -1
	a2i := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	for _, r := range s {
		i, ok := a2i[string(r)]
		if !ok {
			continue
		}

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
