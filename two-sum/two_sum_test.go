package main

import (
	"reflect"
	"testing"
)

type twoSumTestCase struct {
	name     string
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	testCases := []twoSumTestCase{
		{
			name:     "empty slice",
			nums:     []int{},
			target:   0,
			expected: []int{},
		},
		{
			name:     "one element",
			nums:     []int{1},
			target:   1,
			expected: []int{},
		},
		{
			name:     "two elements",
			nums:     []int{1, 2},
			target:   3,
			expected: []int{0, 1},
		},
		{
			name:     "no solution",
			nums:     []int{1, 2, 3, 4, 5},
			target:   10,
			expected: []int{},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -3,
			expected: []int{0, 1},
		},
		{
			name:     "mix of positive and negative numbers",
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "target is zero",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
