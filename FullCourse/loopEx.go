package main

import "fmt"

func loopEx() {

	// arr := [5]int{2, 5, 3, 1, 6}
	// for i, val := range arr {
	// 	fmt.Printf("indices: %d, value: %d\n", i, val)
	// }

	// arr := []int{1, 2, 3, 2, 12, 4}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("arr[%d]: %d\n", i, arr[i])
	// }

	// arr := []int{1, 1, 1, 1, 1, 1, 1, 1}
	// for i, val := range arr {
	// 	// val += i //by copy

	// 	val++       //to cancel not used err.
	// 	arr[i] += i //by ref
	// }
	// fmt.Println(arr)

	// c := time.After(5 * time.Second) //send a signal after 5 seconds

	// for { //infinity loop

	// 	b := false
	// 	select {
	// 	case <-c: //this line execute when signal received.
	// 		b = true
	// 	default:
	// 		fmt.Println(time.Now())
	// 		time.Sleep(1 * time.Second)
	// 	}
	// 	if b {
	// 		break
	// 	}
	// }

	// //get index and value from arr
	// animals := []string{"dog", "cat", "pig", "bird"}
	// for i, item := range animals {
	// 	fmt.Println(i, ": ", item)
	// }

	//get index and value from map
	myMap := make(map[string]string)
	myMap["dog"] = "this is dog"
	myMap["cat"] = "this is cat"
	myMap["bird"] = "this is bird"
	for index, value := range myMap {
		fmt.Println(index, ": ", value)
	}

}
