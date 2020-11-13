// Problem: https://leetcode.com/problems/self-dividing-numbers/
package main

func selfDividingNumbers(left int, right int) (output []int) {
	for i := left; i <= right; i++ {
		if isSelfDividingNum(i) {
			output = append(output, i)
		}
	}

	return
}

func isSelfDividingNum(num int) bool {
	for i := num; i > 0; i /= 10 {
		if i%10 == 0 || num%(i%10) != 0 {
			return false
		}
	}

	return true
}
