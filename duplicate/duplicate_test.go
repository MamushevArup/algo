package main

import (
	"reflect"
	"testing"
)

type removeDuplicateTestCase struct {
	name     string
	input    []int
	expected []int
}

func TestRemoveDuplicate(t *testing.T) {
	testCases := []removeDuplicateTestCase{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "one element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all duplicates",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "some duplicates",
			input:    []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "negative dups",
			input:    []int{-44, -454, -232, -323, -232, -654, -856, -900, -1, -0, -1, -1},
			expected: []int{-44, -454, -232, -323, -654, -856, -900, -1, 0},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := removeDuplicate(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
