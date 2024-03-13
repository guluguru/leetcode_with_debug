package longest_substring_without_repeating_charactersc

func lengthOfLongestSubstring(s string) int {
	d := make(map[rune]int)
	l, res := 0, 0
	for r, v := range s {
		if _, e := d[v]; !e {
			d[v] = r
		} else {
			lastPresent := d[v]
			for i := l; i < lastPresent; i++ {
				delete(d, rune(s[i]))
			}
			l = lastPresent + 1
			d[v] = r
		}
		res = max(res, r-l+1)
	}
	return res
}
