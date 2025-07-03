package main

import "math"

// time complexity: O(n^2)
func brute_max_sub_array(nums []int) int {
	ms := math.MinInt32
	for i := 0; i < len(nums); i++ {
		cs := 0
		for j := i; j < len(nums); j++ {
			cs += nums[j]
			ms = max(ms, cs)
		}
	}
	return ms
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max_sub_array(nums []int) int {
	max_val, cur_val := math.MinInt64, 0
	for i := 0; i < len(nums); i++ {
		cur_val = max(nums[i], cur_val+nums[i])
		max_val = max(cur_val, max_val)
	}
	return max_val
}
