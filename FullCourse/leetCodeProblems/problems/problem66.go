package problems

//https://leetcode.com/problems/plus-one/

func plusOne(digits []int) []int {
	sol := []int{0}
	sol = append(sol, digits...)
	for i := len(sol) - 1; i >= 0; i-- {
		sol[i]++
		if sol[i] > 9 {
			sol[i] = 0
			continue
		}
		break
	}
	if sol[0] == 0 {
		sol = sol[1:len(sol)]
	}
	return sol
}
