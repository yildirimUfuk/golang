package main

import "fmt"

type myStruct struct {
	name      string
	surname   string
	structArr []int
}

func receiversEx() {
	var myObj myStruct
	first, second := myObj.myStructInitializer("ufuk", "yildirim")

	fmt.Println(myObj)
	fmt.Println(first, second)
}
func (k *myStruct) myStructInitializer(name string, surname string) (string, string) {
	k.name = name
	k.surname = surname
	return "first param", "second param"
}
