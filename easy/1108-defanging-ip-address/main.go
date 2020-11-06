// Problem: https://leetcode.com/problems/defanging-an-ip-address/
package main

import (
	"strings"
)

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1) // standard library, cleaner solution

	/* my own solution, probably dirtier xP
	defanged := make([]rune, 0)

	for i := 0; i < len(address); i++ {
		temp := string(address[i])
		if string(address[i]) == "." {
			temp = "[.]"
		}

			defanged = append(defanged, []rune(temp)...)
		}

	return string(defanged)
	*/
}
