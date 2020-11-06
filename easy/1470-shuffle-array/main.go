// Problem: https://leetcode.com/problems/shuffle-the-array/
package main

import "fmt"

func shuffle(nums []int, n int) []int {
	output := make([]int, 0, len(nums))

	for i := 0; i < n; i++ {
		output = append(output, nums[i])
		output = append(output, nums[i+n])
	}

	return output
}

func main() {
	output := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	fmt.Println(output)
}
