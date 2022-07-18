package slidingwindow

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxProfit(prices []int) int {
	maxProfit := 0
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit = profit + prices[i] - prices[i-1]
		profit = Max(profit, 0)
		maxProfit = Max(maxProfit, profit)
	}
	return maxProfit
}
