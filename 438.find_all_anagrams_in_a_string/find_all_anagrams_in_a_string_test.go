package find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		s      string
		p      string
		expect []int
	}{
		{
			s:      "cbaebabacd",
			p:      "abc",
			expect: []int{0, 6},
		},
		{
			s:      "abab",
			p:      "ab",
			expect: []int{0, 1, 2},
		},
		// Add more test cases here
	}

	for _, test := range tests {
		result := findAnagrams(test.s, test.p)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("Test case failed: s=%s, p=%s, expect=%v, got=%v", test.s, test.p, test.expect, result)
		}
	}
}
