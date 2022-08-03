package leetcode

// 初手3重ループで出したらさすがにTime Limit Exceededだったので
// 二分探索で

func mySqrt(x int) int {
	min := 1
	max := x / 2
	for min <= max {
		mid := (min + max) / 2
		pow := mid * mid

		if pow <= x {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return min - 1
}
