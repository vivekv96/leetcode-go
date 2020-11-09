// Problem: https://leetcode.com/problems/minimum-time-visiting-all-points/
package main

func minTimeToVisitAllPoints(points [][]int) int {
	time := 0

	for i := 0; i < len(points)-1; i++ {
		a := points[i]
		b := points[i+1]

		x1, y1, x2, y2 := a[0], a[1], b[0], b[1]

		max := abs(x2 - x1)
		if d := abs(y2 - y1); d > max {
			max = d
		}

		time += max
	}

	return time
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
