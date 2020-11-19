// Problem: https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
package main

func repeatedNTimes(A []int) int {
	m := make(map[int]int)
	output := 0

	for i := range A {
		m[A[i]] += 1

		if m[A[i]] > 1 {
			output = A[i]
			break
		}
	}

	return output
}
