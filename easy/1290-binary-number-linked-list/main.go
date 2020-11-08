// Problem: https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	output := 0

	for head != nil {
		output *= 2
		output += head.Val

		head = head.Next
	}

	return output
}

/* Naïve approach
func getDecimalValue(head *ListNode) int {
	if head.Val == 0 && head.Next == nil {
		return 0
	}

	output := 0
	n := 1
	nums := make([]int, 0)

	nums = append(nums, head.Val)

	for head.Next != nil {
		n++
		head = head.Next
		nums = append(nums, head.Val)
	}

	for i, num := range nums {
		if num == 1 {
			output += int(math.Pow(2, float64(n-i-1)))
		}
	}

	return output
}
*/
