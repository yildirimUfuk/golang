package main

import "fmt"

func funcEx(str string, arr ...string) {
	//defers run after a function return
	defer fmt.Println("first defer line")
	defer fmt.Println("second defer line") //last one execute first

	fmt.Printf("first param: %s, arr first param: %s, arr second param: %s\n", str, arr[0], arr[1])
	return
}
