// Problem: https://leetcode.com/problems/number-of-good-pairs/
package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	numOfPairs := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				numOfPairs++
			}
		}
	}

	return numOfPairs
}

func main() {
	output := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	fmt.Println(output)
}
