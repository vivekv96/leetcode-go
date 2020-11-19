// Problem: https://leetcode.com/problems/robot-return-to-origin/
package main

func judgeCircle(moves string) bool {
	var x, y int

	for _, s := range moves {
		switch s {
		case 'L':
			x--
		case 'R':
			x++
		case 'U':
			y++
		case 'D':
			y--
		}
	}

	return x == 0 && y == 0
}
