// Problem: https://leetcode.com/problems/split-a-string-in-balanced-strings/
package main

func balancedStringSplit(s string) int {
	balanced := 0
	count := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "L" {
			balanced++
		}

		if string(s[i]) == "R" {
			balanced--
		}

		if balanced == 0 {
			count++
		}
	}

	return count
}
