package main

// twoSum find the indices of the two numbers such that they add up to a specific target
// in map I store the value and index of the value for O(1) access which boost the performance
// Time complexity is O(n)
// Space complexity is O(n)
// where n is the length of the slice
func twoSum(nums []int, target int) []int {
	hm := make(map[int]int, len(nums))

	for i, v := range nums {
		diff := target - v
		if value, ok := hm[diff]; ok && i != value {
			return []int{hm[diff], i}
		}
		hm[v] = i
	}
	return []int{}
}
