// Problem: https://leetcode.com/problems/running-sum-of-1d-array/
package main

import "fmt"

func runningSum(nums []int) []int {
	runningSum := make([]int, len(nums))

	runningSum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		runningSum[i] = runningSum[i-1] + nums[i]
	}

	return runningSum
}

func main() {
	output := runningSum([]int{1, 2, 3, 4})
	fmt.Println(output)
}
