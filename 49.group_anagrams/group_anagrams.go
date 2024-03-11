package group_anagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	var h = make(map[string][]string)
	for _, s := range strs {
		ns := sortString(s)
		if _, ok := h[ns]; ok {
			h[ns] = append(h[ns], s)
		} else {
			h[ns] = []string{s}
		}
	}
	res := make([][]string, 0, len(h))
	for _, v := range h {
		res = append(res, v)
	}
	return res
}

func sortString(s string) string {
	var a = []byte(s)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return string(a)
}
