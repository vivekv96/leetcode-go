// Problem: https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
package main

func countNegatives(grid [][]int) (count int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				count += len(grid[i]) - j
				break
			}
		}
	}

	return
}
