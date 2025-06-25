package main

func brute_max_profit_day(prices []int) int {
	//brute force
	max_profit := 0
	for i:=0; i<len(prices); i++{
		for j:=i+1; j<len(prices); j++ {
			diff := prices[j] - prices[i]
			if diff > max_profit{
				max_profit = diff
			}
		}
	}
	return max_profit
}

func max_profit_day(prices []int) int {
	//brute force
	max_profit := 0
	min_profit := prices[0]
	for i:=0; i<len(prices); i++ {
		if prices[i] < min_profit{
			min_profit = prices[i]
		} else if prices[i] - min_profit > max_profit{
			max_profit = prices[i] - min_profit
		}
	}
	return max_profit
}