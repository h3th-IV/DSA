package main

import "fmt"

func main(){
	prices := []int{7, 1, 5, 3, 6, 4}

	brute_max_profit := brute_max_profit_day(prices)
	fmt.Println(brute_max_profit)

	max_profit := max_profit_day(prices)
	fmt.Println(max_profit)
}