package main

import "fmt"

func pointerEx() {
	type String *string

	str := "this is the string"
	var p String = &str
	*p = "asdf"

	fmt.Println(str, " - ", *p)
}
