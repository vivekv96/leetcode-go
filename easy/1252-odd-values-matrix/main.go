// Problem: https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
package main

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)

	for r := 0; r < n; r++ {
		matrix[r] = make([]int, m)

		for c := 0; c < m; c++ {
			matrix[r][c] = 0
		}
	}

	for i := 0; i < len(indices); i++ {
		for j := 0; j < 2; j++ {
			index := indices[i][j]

			if j == 0 {
				for c := 0; c < m; c++ {
					matrix[index][c] += 1
				}
			} else {
				for r := 0; r < n; r++ {
					matrix[r][index] += 1
				}
			}
		}
	}

	count := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if matrix[r][c]%2 == 1 {
				count++
			}
		}
	}

	return count
}
