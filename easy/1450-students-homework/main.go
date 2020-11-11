// Problem: https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/
package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0

	for i := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			count++
		}
	}

	return count
}
