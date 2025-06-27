package main

import "fmt"

func brute_two_sum(nums []int, target int) []int{
	for i:=0; i<len(nums); i++{
		for j:=0; j<len(nums)-1; j++{
			if nums[i] + nums[j] == target{
				sum := nums[i] + nums[j]
				fmt.Println(sum)
				return []int{i, j}
			}
		}
	}
	return []int{0,0}
}

func two_sum(nums []int, target int) []int{
	
}