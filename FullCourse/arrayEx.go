package main

import "fmt"

func arrayExFunc() {

	var arr [10]int
	arr[0] = 1
	fmt.Printf("var arr [10]int => arr[0]= %d\n", arr[0])

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("var arr2=[5]int{1,2,3,4,5} => arr2=%d\n", arr2)

	arr3 := make([]int, 3)
	arr3[0] = 1
	println(arr3[2])

	arr4 := []int{1, 23, 21, 31, 23, 1, 2}
	fmt.Println(arr4)

	arr5 := [...]int{1, 23, 1, 23, 12, 31, 23, 123, 1}
	fmt.Println(arr5)

}
