package move_zeroes

func moveZeroes(nums []int) {
	l, r := 0, 0
	n := len(nums)
	for r != n {
		if nums[l] == 0 {
			if nums[r] == 0 {
				r++
			} else {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r++
			}
		} else {
			l++
			r++
		}
	}
}
