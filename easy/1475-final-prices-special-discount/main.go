// Problem: https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/
package main

func finalPrices(prices []int) []int {
	output := make([]int, 0, len(prices))

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				output = append(output, prices[i]-prices[j])
				break
			}
		}

		if len(output)-1 != i {
			output = append(output, prices[i])
		}
	}

	return output
}
