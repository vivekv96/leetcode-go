// Problem: https://leetcode.com/problems/xor-operation-in-an-array/
package main

func xorOperation(n int, start int) int {
	nums := make([]int, n)
	output := 0

	for i := 0; i < n; i++ {
		nums[i] = start + 2*i
		output ^= nums[i]
	}

	return output
}
