// Problem: https://leetcode.com/problems/di-string-match/
package main

func diStringMatch(S string) []int {
	l, r := 0, len(S)

	output := make([]int, len(S)+1)

	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			output[i] = l
			l++
		} else {
			output[i] = r
			r--
		}
	}

	output[len(S)] = l

	return output
}
