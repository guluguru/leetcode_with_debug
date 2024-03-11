package move_zeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{0, 0, 0, 1, 0}, []int{1, 0, 0, 0, 0}},
	}

	for _, test := range tests {
		input := make([]int, len(test.nums))
		copy(input, test.nums)

		moveZeroes(input)

		if !reflect.DeepEqual(input, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, input)
		}
	}
}
