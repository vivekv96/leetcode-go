// Problem: https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
package main

func sumZero(n int) []int {
	output := make([]int, 0, n)

	for i := n; i > 1; i-- {
		output = append(output, -i)
	}

	output = append(output, (n*(n+1))/2-1)

	return output
}
