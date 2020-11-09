// Problem: https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
package main

func findNumbers(nums []int) int {
	count := 0

	for _, num := range nums {
		l := 0
		for num != 0 {
			num /= 10
			l++
		}

		if l%2 == 0 {
			count++
		}
	}

	return count
}

/* Another approach -
import "strconv"

func findNumbers(nums []int) (count int) {
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			count++
		}
	}
	return
}
*/
