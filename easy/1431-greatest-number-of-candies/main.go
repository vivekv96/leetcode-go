// Problem: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	output := make([]bool, len(candies))
	max := candies[0]

	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			output[i] = true
		} else {
			output[i] = false
		}
	}

	return output
}
