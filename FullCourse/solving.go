package main

import (
	"strconv"
)

//leetcode 1. question.
func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i, val := range nums {
		if num, ok := myMap[target-val]; ok {
			return []int{num, i}
		}
		myMap[val] = i
	}
	return []int{0, 0}
}

//leetcode 9. question
func isPalindrome(x int) bool {
	strNumber := strconv.Itoa(x)
	for i := 0; i < len(strNumber)/2; i++ {
		if strNumber[i] != strNumber[len(strNumber)-i-1] {
			return false
		}
	}
	return true
}
