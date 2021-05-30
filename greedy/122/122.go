package greedy

/*
title: Best Time to Buy and Sell Stock II
conent：
You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete as many transactions as you like
 (i.e., buy one and sell one share of the stock multiple times).
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the
stock before you buy again).
Example：
Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.


解题思路：
*/
func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		if i < len(prices)-1 && prices[i] < prices[i+1] {
			profit = profit + prices[i+1] - prices[i]
		}

	}
	return profit
}
