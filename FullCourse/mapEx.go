package main

import "fmt"

func mapEx() {
	var map1 map[int]string
	map1 = make(map[int]string) // it must initialize by 'make'.
	map1[0] = "abuzer"
	map1[1] = "kamil"
	map1[4] = "asdf"
	fmt.Printf("map[0]: %s, map[1]: %s map[4]: %s", map1[0], map1[1], map1[3])

	map2 := make(map[int]string)
	map2[0] = "zero"
	map2[4] = "four"
	fmt.Printf("map2[0]: %s, map2[4]: %s", map2[0], map2[4])

	map3 := map[string]int{
		"abuzer":      1,
		"kamil":       2,
		"abdulcabbar": 3,
	}
	fmt.Println(map3)

	if val, ok := map3["kamil"]; ok { //if key(kamil) exist ok becomes true.
		fmt.Println("valuse is : ", val)
	}
	delete(map3, "kamil")
	if val, ok := map3["kamil"]; ok {
		fmt.Print(val)
	} else {
		fmt.Println("can not found key!")
	}

	map4 := map[string][]string{
		// "abuzer": []string{"first", "second", "thirt"}, //redundant type warning
		"abuzer": {"first", "second", "thirt"},
		"kamil":  {"first", "second", "thirt"},
	}
	fmt.Println(map4)

	map5 := map[int]string{
		0: "0000000",
		1: "1111111",
	}
	for k, v := range map5 {
		fmt.Println("->", k, v)
	}

}
