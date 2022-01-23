package main

import (
	"fmt"
	"myPackage"
	"sync"
	"time"
)

var wgChannel sync.WaitGroup

func chanFunc(myChannel chan string) {
	time.Sleep(time.Second * 3)
	myChannel <- "channelFunc message" //message send to channel
}
func channelsEx() {
	// var myChannel chan string
	// myChannel = make(chan string)
	myChannel := make(chan string)

	go chanFunc(myChannel)
	fmt.Println("myChannel:", myChannel)     //addres of channel
	fmt.Println("<-myChannel:", <-myChannel) //value of channel

}

// ////////////////////////////////////////////////////
// //?? SECOND EXAMPLE! ??//
// func routine1(count *int) {
// 	for i := 0; i < 100000000; i++ {
// 		*count++
// 	}
// 	wgChannel.Done()
// }
// func routine2(count *int) {
// 	for i := 0; i < 10000000; i++ {
// 		*count++
// 	}
// 	wgChannel.Done()
// }

// func channelsEx2() {
// 	wgChannel.Add(2)
// 	count := 0

// 	go routine1(&count)
// 	go routine2(&count)

// 	wgChannel.Wait()
// 	fmt.Println("without channel val of i: ", count)
// }

// ////////////////////////////////////////////////
// //?? THIRT EXAMPLE ??//
// func firstRoutine(num chan int) {
// 	i := 0
// 	time.Sleep(time.Second * 2)
// 	for ; i < 1000000000; i++ {
// 	}
// 	fmt.Println("first routine end")
// 	num <- i
// }
// func secondRoutine(num2 chan int) {
// 	i := 0
// 	for ; i < 1000000000; i++ {
// 	}
// 	fmt.Println("second routine end")
// 	num2 <- i
// }
// func thirtRotuine() {
// 	time.Sleep(time.Second * 10)
// 	fmt.Println("thirt routine end") //it doesn't guarantied this line will be axecuted.
// }
// func channelsEx3() {
// 	defer func() { // if it throws panic this will catche what the err is.
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	num := make(chan int)
// 	num2 := make(chan int)

// 	fmt.Println("first routine calling with channel")
// 	go firstRoutine(num)
// 	fmt.Println("num: ", <-num)

// 	fmt.Println("second routine calling with channel")
// 	go secondRoutine(num2)
// 	fmt.Println("num2: ", <-num2)

// 	fmt.Println("thirt rotuine calling without channel")
// 	go thirtRotuine()
// }

////////////////////////////////////////////////
//?? THIRT EXAMPLE ??//
const maxNum = 100

func channelsEx4() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateVal(intChan)
	fmt.Println(<-intChan)
}
func calculateVal(intChan chan int) {
	// randomNumber := myPackage.RandomNumber(maxNum)
	// intChan <- randomNumber

	myPackage.RandomNumber2(intChan, maxNum)
	fmt.Println(<-intChan)
}
