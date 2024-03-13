package trapping_rain_water

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{4, 0, 11, 0, 11, 5, 1, 4}, 18},
		{[]int{4, 0, 3, 0, 11, 5, 1, 4}, 12},
		{[]int{4, 0, 3, 0, 11, 5, 1, 3}, 11},
	}

	for _, test := range tests {
		got := trap(test.height)
		if got != test.want {
			t.Errorf("trap(%v) = %d; want %d", test.height, got, test.want)
		}
	}
}
