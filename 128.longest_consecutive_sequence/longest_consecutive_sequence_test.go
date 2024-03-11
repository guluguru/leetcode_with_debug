package longest_consecutive_sequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
