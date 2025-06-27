package main

import "fmt"

func main(){
	prices := []int{7, 1, 5, 3, 6, 4}

	brute_max_profit := brute_max_profit_day(prices)
	fmt.Println(brute_max_profit)

	max_profit := max_profit_day(prices)
	fmt.Println(max_profit)

	nums := []int{2, 7, 4, 5, 8}
	brute_axis := brute_two_sum(nums, 9)
	fmt.Println(brute_axis)

	axis := two_sum(nums, 9)
	fmt.Println(axis)
}