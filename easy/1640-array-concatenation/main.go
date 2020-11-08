// Problem: https://leetcode.com/problems/check-array-formation-through-concatenation/
package main

func canFormArray(arr []int, pieces [][]int) bool {
	p := make(map[int][]int)

	for i, piece := range pieces {
		p[pieces[i][0]] = piece[1:]
	}

	i := 0
	for i < len(arr) {
		slc, ok := p[arr[i]]
		if !ok {
			return false
		}

		i++
		for _, ele := range slc {
			if ele != arr[i] {
				return false
			}

			i++
		}
	}

	return true
}
