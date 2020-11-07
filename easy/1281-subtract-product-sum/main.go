// Problem: https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
package main

func subtractProductAndSum(n int) int {
	sum, product := 0, 1

	for {
		if n == 0 {
			break
		}

		last := n % 10
		n /= 10

		sum += last
		product *= last
	}

	return product - sum
}
