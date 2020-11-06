// Problem: https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
package main

func numberOfSteps(num int) int {
	count := 0

	for {
		if num == 0 {
			return count
		}

		count++

		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
	}
}
