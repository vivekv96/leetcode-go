// Problem: https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
package main

func smallerNumbersThanCurrent(nums []int) []int {
	output := make([]int, len(nums), len(nums))

	for i, _ := range nums {
		count := 0

		for j, _ := range nums {
			if nums[j] < nums[i] {
				count++
			}
		}

		output[i] = count
	}

	return output
}
