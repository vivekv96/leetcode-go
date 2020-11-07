// Problem: https://leetcode.com/problems/create-target-array-in-the-given-order/
package main

func createTargetArray(nums []int, index []int) []int {
	target := make([]*int, len(nums), len(nums))

	for i, num := range nums {
		if target[index[i]] != nil {
			copy(target[index[i]+1:], target[index[i]:])
		}

		temp := num
		target[index[i]] = &temp
	}

	output := make([]int, 0)

	for _, num := range target {
		output = append(output, *num)
	}

	return output
}
