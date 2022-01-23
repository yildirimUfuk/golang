package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //waitgrup variable.

func goRoutines() {
	i := 0
	j := 0
	wg.Add(2) //this group has 2 go routines.
	go goRoutines1(i)
	go goRoutines2(j)
	normalFunc()
	wg.Wait() // this guarantied 2 routines must be done. but it doesn't guarantied which one will be compleated first.

}
func normalFunc() {
	println("normal function executed!")
}

func goRoutines1(i int) {
	println("beginning goRoutines1")
	for i < 99999999 {
		i++
	}
	fmt.Println("goRoutine1: ", i)
	wg.Done() // one routine adds to wg.
}

func goRoutines2(i int) {
	println("beginning goRoutines2")
	for i < 100000000 {
		i++
	}
	fmt.Println("goRoutine 2: ", i)
	wg.Done() // one routine adds to wg
}
