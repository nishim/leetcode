package leetcode

// https://leetcode.com/problems/container-with-most-water/submissions/

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j, h := range height[i:] {
			var low int
			if height[i] > h {
				low = h
			} else {
				low = height[i]
			}

			if max < low*j {
				max = low * j
			}
		}
	}
	return max
}
