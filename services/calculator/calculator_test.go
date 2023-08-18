package calculator

import (
	"reflect"
	"testing"
)

func TestFindExactSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{250, 500, 1000, 2000, 5000},
			target:   250,
			expected: []int{250},
		},
		{
			nums:     []int{250, 500, 1000, 2000, 5000},
			target:   251,
			expected: []int{500},
		},
		{
			nums:     []int{250, 500, 1000, 2000, 5000},
			target:   501,
			expected: []int{500, 250},
		},
		{
			nums:     []int{250, 500, 1000, 2000, 5000},
			target:   12001,
			expected: []int{5000, 5000, 2000, 250},
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			target:   8,
			expected: []int{5, 3},
		},
		{
			nums:     []int{2, 4, 6, 8, 10},
			target:   16,
			expected: []int{10, 6},
		},
		{
			nums:     []int{10, 20, 30},
			target:   45,
			expected: []int{30, 20},
		},
	}

	for _, test := range tests {
		result := FindPerfectPackets(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums %v and target %d, expected %v, but got %v", test.nums, test.target, test.expected, result)
		}
	}
}
