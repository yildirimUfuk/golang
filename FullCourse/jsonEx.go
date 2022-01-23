package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	HairColor string `json:"hairColor"`
	HasAnimal bool   `json:"hasAnimal"`
}

func jsonEx1() {
	jsonToStruct()
	structToJson()

}

//it converts json obj to structs
func jsonToStruct() {
	defer func() {
		if err := recover(); err != nil {
			println("deger err: ", err)
		}
	}()

	myJson :=
		`[{"firstName":"Abuzer",
			"lastName":"kamil",
			"hairColor":"black",
			"hasAnimal":true},

		{"firstName":"Abuzer2",
			"lastName":"Kamil2",
			"hairColor":"pink",
			"hasAnimal":false}]`
	var unmarshalled []person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(unmarshalled)
}

//it converts struct to json obj
func structToJson() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("struct to json err: ", err)
		}
	}()
	var p1 person
	p1.FirstName = "abuzer"
	p1.LastName = "kamil"
	p1.HairColor = "red"
	p1.HasAnimal = false

	var p2 person
	p2.FirstName = "abuzer2"
	p2.LastName = "kamil2"
	p2.HairColor = "black"
	p2.HasAnimal = true

	mySlice := make([]person, 0)
	mySlice = append(mySlice, p1)
	mySlice = append(mySlice, p2)

	newJson, err := json.MarshalIndent(mySlice, "", "	")
	if err != nil {
		fmt.Println("marshalIndent err!")
	}
	fmt.Println(string(newJson))
}
