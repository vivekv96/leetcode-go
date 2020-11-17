// Problem: https://leetcode.com/problems/sort-array-by-parity/
package main

func sortArrayByParity(A []int) []int {
	even := make([]int, 0)
	odd := make([]int, 0)

	for i := range A {
		if A[i]%2 == 0 {
			even = append(even, A[i])
			continue
		}

		odd = append(odd, A[i])
	}

	return append(even, odd...)
}
