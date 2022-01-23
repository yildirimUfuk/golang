package main

import "fmt"

//slices example begin
func slicesEx() {
	// var slc []int
	// slc = append(slc, 1)
	// fmt.Printf("slc: %d, slc length: %d, slc cap: %d", slc, len(slc), cap(slc))

	// slc2 := make([]int, 0, 3) // if length less than 3 cap will be 3. if 4th element add capacity becomes 6. if length become 7 it becomes 12. it decrease *2 by thirth parameter of make func.
	// slc2 = append(slc2, 1)
	// fmt.Printf("slc2: %d, slc2 length: %d, slc2 cap: %d", slc2, len(slc2), cap(slc2))

	// slc1 := []int{1, 2, 3, 4, 5}
	// slc1 = append(slc1, 6)
	// slc2 := []int{7, 8}
	// // slc3 := append(slc1, slc2) //error !
	// slc3 := append(slc1, slc2...) // it add slc2's int values to slc1.
	// fmt.Println(slc3)

	var slc4 []int
	fmt.Printf("%#v", slc4) //slc default value "nil"

}
