package main

import "fmt"

func main(){
	// prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{9,2,5,3,6,7,1}


	brute_max_profit := brute_max_profit_day(prices)
	fmt.Println(brute_max_profit)

	max_profit := max_profit_day(prices)
	fmt.Println(max_profit)

	nums := []int{2, 7, 4, 5, 8}
	brute_axis := brute_two_sum(nums, 9)
	fmt.Println(brute_axis)

	axis := two_sum(nums, 9)
	fmt.Println(axis)

	dup_array := []int{1,1,2,3,4,4}
	dup_exist := check_duplicate(dup_array)
	fmt.Println(dup_exist)
}