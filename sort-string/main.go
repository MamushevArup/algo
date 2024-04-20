package main

// bubbleSort sorts a slice of runes using bubble sort algorithm
// regarding the instruction is it not clear will we get only english lowercase letters or not
// so I use rune type instead of byte to support all unicode characters
// since I am not allowed to use any built-in sort functions and data structures
// I implement bubble sort which not requires any additional memory space,
// but it is not efficient for large slice
// I use one bool variable to check if the slice is already sorted if it is time complexity will be O(n)
// In average and worst case time complexity is O(n^2)
// Space complexity is O(1)
// where n is the length of the slice
func bubbleSort(s []rune) []rune {
	n := len(s)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return s
}
