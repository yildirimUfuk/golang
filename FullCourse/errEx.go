package main

import (
	"errors"
	"fmt"
	"log"
)

func errorEx() {
	if err := createErr(); err != nil {
		log.Fatal(err) //function will be terminated
	}
	fmt.Println("no error")
}
func createErr() error {
	// i := 0 // no error
	i := 1 //error
	if i == 1 {
		return errors.New("error created succecfuly")
	}
	return nil
}
