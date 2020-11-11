// Problem: https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
package main

import (
	"sort"
)

func maxProduct(nums []int) (product int) {
	if len(nums) == 2 {
		return (nums[0] - 1) * (nums[1] - 1)
	}

	sort.Ints(nums)

	product = (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)

	return
}
