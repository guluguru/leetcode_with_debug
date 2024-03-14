package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	// 滑动窗口，维护子串
	// 窗口长度不够，滑动窗口右边界
	// 窗口长度相等，判断是否为子串
	// 窗口长度超过，滑动窗口左边界
	var res []int
	hs := make(map[rune]int)
	hp := make(map[rune]int)
	for _, v := range p {
		hp[v]++
	}
	for i, v := range s {
		hs[v]++
		if i >= len(p)-1 {
			if isAnagram(hp, s[i-len(p)+1:i+1]) { // 能过，但是耗时击败5%，估计是这一步耗时太长
				res = append(res, i-len(p)+1)
			}
			hs[rune(s[i-len(p)+1])]--
		}
	}
	return res
}

func isAnagram(hp map[rune]int, ss string) bool {
	hss := make(map[rune]int)
	for _, v := range ss {
		hss[v]++
	}
	for k, v := range hp {
		if hss[k] != v {
			return false
		}
	}
	return true
}
