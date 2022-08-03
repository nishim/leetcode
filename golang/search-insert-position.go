package leetcode

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	i := 0

	if nums[r] < target {
		return r + 1
	}
	if target == 0 {
		return 0
	}

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid - 1
			i = mid
		} else {
			l = mid + 1
		}
	}

	return i
}
