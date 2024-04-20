package main

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
