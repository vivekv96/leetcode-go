// Problem: https://leetcode.com/problems/defanging-an-ip-address/
package main

import "fmt"

func defangIPaddr(address string) string {
	defanged := make([]rune, 0)

	for i := 0; i < len(address); i++ {
		temp := string(address[i])
		if string(address[i]) == "." {
			temp = "[.]"
		}

		defanged = append(defanged, []rune(temp)...)
	}

	return string(defanged)
}

func main() {
	output := defangIPaddr("1.1.1.1")
	fmt.Println(output)
}
