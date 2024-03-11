package three_sum

import "sort"

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		if nums[r] < 0 || nums[i] > 0 { // 提前终止
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				n2 := nums[l]
				n3 := nums[r]
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
