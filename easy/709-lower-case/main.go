// Problem: https://leetcode.com/problems/to-lower-case/
package main

func toLowerCase(str string) string {
	/* Obvious Standard Library solution -
	return strings.ToLower(str)
	*/

	output := make([]byte, 0, len(str))

	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			c += 32
		}

		output = append(output, byte(c))
	}

	return string(output)
}
