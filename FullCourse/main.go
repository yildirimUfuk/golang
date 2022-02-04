package main

import "fmt"

func main() {
	// web.HttpMain()
	digits := []int{1, 2, 1, 4, 3}

	fmt.Println(plusOne(digits))
}
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

// var portNumber = ":8080"

// //main http package example function.
// func HttpMain() {
// 	http.HandleFunc("/", handlers.Home)
// 	http.HandleFunc("/about", handlers.About)
// 	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))
// 	// fmt.Printf(fmt.Sprintf("starting application on port: %s\n", portNumber))
// 	_ = http.ListenAndServe(portNumber, nil)
// }
