package problems

import "sort"

func threeSum(nums []int) [][]int {

	solution := make([][]int, 1)
	s := make([]int, 1)

	solution = append(solution, s)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if find := sort.SearchInts(nums, -1*(nums[i]+nums[j]));  {

			}
		}
	}

	return solution
}
