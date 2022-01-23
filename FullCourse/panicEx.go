package main

import (
	"fmt"
	"time"
)

func panicEx() {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		if i == 2 {
			panic("unexpected situation handled")
		}
		fmt.Printf("loop i: %d\n", i)
	}
}
