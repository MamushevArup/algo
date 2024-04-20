package main

// longestIncreasingSubsequence find the length of the longest increasing subsequence
// it uses dynamic programming to find the result
// where we should iterate forward and backward for each element and find max value in subsequence
// going from the end to the slice
// Time complexity is O(n^2)
// Space complexity is O(n)
// where n is the length of the slice
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
