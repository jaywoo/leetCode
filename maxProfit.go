//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
package main

import (
	"fmt"
)

func main() {
	prices := []int{7,6,4,3,1}
	j := 0
	price := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[j] {
			price = price + prices[i] - prices[j]
		}
		j++
	}

	fmt.Println(price)
}

// func maxProfit(prices []int) int {

// }
