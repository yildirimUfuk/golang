package main

import "strings"

//string example begin
func stringEx() {
	// str := "this is the string variable"
	// str1 := str[:5]
	// str2 := str[len(str)-8:]
	// fmt.Printf("str: %s, str[:5]: %s, str[len(str)-8:]: %s", str, str1, str2)

	// str1 := "adsf"
	// str2 := "asdf"
	// if strings.EqualFold(str1, str2) {
	// 	println("equal")
	// } else {
	// 	println("not equal")
	// }

	// str1 := "this is the string sentences"

	// if strings.HasPrefix(str1, "this") {
	// 	println("it begins prefix 'this'")
	// } else {
	// 	println("it doesn't begin prefix 'this'")
	// }

	str1 := "this is a string"
	if strings.Contains(str1, "is") {
		println("it contains sutbstr: 'is'")
	} else {
		println("it doesn't contains substr: is")
	}
} //string example end
