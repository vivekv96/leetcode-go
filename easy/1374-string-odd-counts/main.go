// Problem: https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/
package main

func generateTheString(n int) string {
	b := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		b = append(b, 'a')
	}

	if n%2 == 0 {
		b[n-1] = 'b'
	}

	return string(b)
}
