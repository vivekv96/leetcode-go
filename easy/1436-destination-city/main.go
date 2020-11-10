// Problem: https://leetcode.com/problems/destination-city/
package main

func destCity(paths [][]string) string {
	m := make(map[string]struct{})

	// populate `m` with all the starting cities
	for i := 0; i < len(paths); i++ {
		m[paths[i][0]] = struct{}{}
	}

	var output string

	// any city that's on the right hand side and doesn't exist in `m` is
	// our answer
	for i := 0; i < len(paths); i++ {
		if _, ok := m[paths[i][1]]; !ok {
			output = paths[i][1]
			break
		}
	}

	return output
}
