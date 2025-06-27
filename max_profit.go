package main

func brute_max_profit_day(prices []int) int {
	max_profit := 0
	for i:=0; i<len(prices); i++{
		for j:=i+1; j<len(prices); j++{
			diff := prices[j] - prices[i]
			if diff > max_profit{
				max_profit = diff
			}
		}
	}
	return max_profit
}

func max_profit_day(prices []int) int {
	max_profit := 0
	min_sale := prices[0]
	for _, ele := range prices{
		diff := ele - min_sale
		if ele < min_sale{
			min_sale = ele
		} else if diff > max_profit{
			max_profit = diff
		}
	}
	return max_profit
}