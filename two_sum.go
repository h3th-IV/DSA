package main


func brute_two_sum(nums []int, target int) []int{
	for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++{
			if nums[i] + nums[j] == target{
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func two_sum(nums []int, target int) []int{
	num_map := make(map[int]int)
	for ind, ele := range nums{
		diff := target - ele
		if val, exist := num_map[diff]; exist{
			return []int{val, ind}
		}
		num_map[ele] = ind
	}
	return []int{0, 0}
}