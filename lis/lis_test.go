package main

import "testing"

func TestLongestIncreasingSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected int
	}{
		{
			name:     "empty",
			slice:    []int{},
			expected: 0,
		},
		{
			name:     "one element",
			slice:    []int{1},
			expected: 1,
		},
		{
			name:     "increasing",
			slice:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "decreasing",
			slice:    []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "equal elem",
			slice:    []int{0, 0, 0, 0, 0, 0, 0},
			expected: 1,
		},
		{
			name:     "negative decreasing",
			slice:    []int{-1, -2, -3, -4, -5},
			expected: 1,
		},
		{
			name:     "negative increasing",
			slice:    []int{-5, -4, -3, -2, -1},
			expected: 5,
		},
		{
			name:     "top down slice",
			slice:    []int{1, 10, 2, 9, 3, 8, 4, 7, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestIncreasingSubsequence(tt.slice)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
