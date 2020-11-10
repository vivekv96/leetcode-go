// Problem: https://leetcode.com/problems/unique-morse-code-words/
package main

func uniqueMorseRepresentations(words []string) int {
	table := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	t := make(map[string]int)

	for i := 0; i < len(words); i++ {
		s := ""
		for _, c := range words[i] {
			s += table[c-'a']
		}

		t[s]++
	}

	return len(t)
}
