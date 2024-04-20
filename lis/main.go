package main

func longestIncreasingSubsequence(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	currentSub := make([]int, len(slice))

	for i := len(slice) - 1; i >= 0; i-- {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				currentSub[i] = max(currentSub[i], currentSub[j]+1)
			}
		}
	}
	return maxElemSlice(currentSub) + 1
}

func maxElemSlice(arr []int) int {
	maxElem := arr[0]
	for _, v := range arr {
		if maxElem < v {
			maxElem = v
		}
	}
	return maxElem
}
