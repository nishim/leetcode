package leetcode

// memo
// GolangのSliceを関数の引数に渡した時の挙動
// https://christina04.hatenablog.com/entry/2017/09/26/190000

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == n {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			n = nums[i]
		}
	}

	return len(nums)
}
