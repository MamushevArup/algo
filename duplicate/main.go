package main

// removeDuplicate eliminate duplicates from a slice of integers
// this program is not use any additional memory space
// but it is a bit slower than it can be with maps
// time complexity is O(n^2)
// space complexity is O(1) (where we not include the slice which coming from input)
// where n is the length of the slice
func removeDuplicate(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				arr = append(arr[:j], arr[j+1:]...)
				j--
			}
		}
	}

	return arr
}
