package main

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
