package two_sum

func twoSum(nums []int, target int) []int {
	h := make(map[int]int)
	for k2, v := range nums {
		if k1, ok := h[target-v]; ok {
			return []int{k1, k2}
		}

		if _, ok := h[v]; !ok {
			h[v] = k2
		}
	}
	return []int{-1, -1}
}
