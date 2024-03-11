package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	h := make(map[int]bool)
	for _, v := range nums {
		h[v] = true
	}

	set := []int{}
	for k, _ := range h {
		set = append(set, k)
	}

	res := 0
	// 遍历，略过v-1存在的，对于不存在的，往后遍历并计算结果
	for _, v := range set {
		if _, exist := h[v-1]; exist {
			continue
		}
		i, cres := 1, 1
		for {
			if _, exist := h[v+i]; !exist {
				break
			} else {
				cres++
				i++
			}
		}
		if cres > res {
			res = cres
		}
	}
	return res
}
