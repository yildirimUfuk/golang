package problems

import "strings"

//https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		count++
	}
	return count
}

func lengthOfLastWord2(s string) int {
	s = strings.TrimSpace(s)
	i := strings.LastIndex(s, " ")
	return len(s) - i - 1
}
