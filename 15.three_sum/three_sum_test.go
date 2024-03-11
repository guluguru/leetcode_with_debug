package three_sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{1, 2, 3, 4, 5}, [][]int{}},
		{[]int{-2, -1, 0, 1, 2, 3}, [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
	}

	for _, test := range tests {
		got := ThreeSum(test.nums)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ThreeSum(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}
