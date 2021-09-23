package data_structure1

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			if tmp := prices[i] - minPrice; tmp > maxProfit {
				maxProfit = prices[i] - minPrice
			}
		}
	}
	return maxProfit
}
