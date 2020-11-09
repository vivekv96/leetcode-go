// Problem: https://leetcode.com/problems/count-good-triplets/
package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0

	/*
		According to the question -
			A triplet (arr[i], arr[j], arr[k]) is good if the following conditions are true:

			0 <= i < j < k < arr.length
			|arr[i] - arr[j]| <= a
			|arr[j] - arr[k]| <= b
			|arr[i] - arr[k]| <= c

		Converting these constraints to code looks like the following -
	*/

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i]-arr[j]) <= a { // This saves a few computations but doesn't change the Time complexity.
				for k := j + 1; k < len(arr); k++ {
					if abs(arr[j]-arr[k]) <= b &&
						abs(arr[i]-arr[k]) <= c {
						count++
					}
				}
			}
		}
	}

	return count
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
