package main

import (
	"fmt"
	"sort"
)

func main() {
	// web.HttpMain()
	digits := []int{1, 2, 5, 7, 9, -10}
	threeSum(digits)

}

func threeSum(nums []int) [][]int {

	type solutionItemStruct struct {
		i, j, k int
	}

	solution := make([][]int, 1)
	digitsMap := make(map[int]int)
	for i, val := range nums {
		digitsMap[val] = i
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			searching := -1 * (nums[i] + nums[j])
			if val, ok := digitsMap[searching]; ok && (val != i && val != j && i != j) {
				s := []int{i, j, val}
				sort.Ints(s)

			}
		}
	}

	for i := 1; i < len(solution); i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(solution[i][0], ':', solution[i][1], ':', solution[i][2])
		}
	}
	return solution
}

// var portNumber = ":8080"

// //main http package example function.
// func HttpMain() {
// 	http.HandleFunc("/", handlers.Home)
// 	http.HandleFunc("/about", handlers.About)
// 	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))
// 	// fmt.Printf(fmt.Sprintf("starting application on port: %s\n", portNumber))
// 	_ = http.ListenAndServe(portNumber, nil)
// }
