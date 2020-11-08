// Problem: https://leetcode.com/problems/shuffle-string/
package main

func restoreString(s string, indices []int) string {
	output := make([]byte, len(s))

	for i := 0; i < len(indices); i++ {
		output[indices[i]] = s[i]
	}

	return string(output)
}
