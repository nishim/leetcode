package leetcode

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	stack := []string{}
	for _, c := range s {
		t := string(c)

		if len(stack) > 0 {
			if (t == ")" && stack[len(stack)-1] == "(") ||
				(t == "}" && stack[len(stack)-1] == "{") ||
				(t == "]" && stack[len(stack)-1] == "[") {
				stack = stack[:len(stack)-1]
				continue
			}
		}

		stack = append(stack, t)
	}

	return len(stack) == 0
}
