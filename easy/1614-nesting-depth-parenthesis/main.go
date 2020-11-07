// Problem: https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
package main

func maxDepth(s string) int {
	max, count := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		}

		if count > max {
			max = count
		}

		if s[i] == ')' {
			count--
		}
	}

	return max
}
