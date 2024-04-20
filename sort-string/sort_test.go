package main

import (
	"reflect"
	"testing"
)

type bubbleSortTestCase struct {
	name     string
	input    string
	expected string
}

func TestBubbleSort(t *testing.T) {
	testCases := []bubbleSortTestCase{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "already sorted",
			input:    "abc",
			expected: "abc",
		},
		{
			name:     "reverse sorted",
			input:    "cba",
			expected: "abc",
		},
		{
			name:     "all same characters",
			input:    "aaa",
			expected: "aaa",
		},
		{
			name:     "random order",
			input:    "bdca",
			expected: "abcd",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			runes := []rune(tt.input)
			result := bubbleSort(runes)
			if !reflect.DeepEqual(string(result), tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, string(result))
			}
		})
	}
}
