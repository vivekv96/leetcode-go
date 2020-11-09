// Problem: https://leetcode.com/problems/maximum-69-number/
package main

import (
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	s := strconv.Itoa(num)

	temp := strings.Replace(s, "6", "9", 1)

	output, _ := strconv.Atoi(temp)
	return output
}
