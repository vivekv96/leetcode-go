// Problem: https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
package main

func replaceElements(arr []int) []int {
	max := 0

	for i := len(arr) - 1; i >= 0; i-- {
		temp := max

		if arr[i] > max {
			max = arr[i]
		}

		arr[i] = temp
	}

	arr[len(arr)-1] = -1

	return arr
}
