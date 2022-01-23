package main

import "fmt"

// to interface syntax: type variableName interface{}
// if a struct inherit an interface, the struct must implement receiver all function of interface.
// a struct can have much more receiver except

func interfaceExfunc() {
	dog := dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}
	gorilla := gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	printInfo(&dog)
	printInfo(&gorilla)

}

func printInfo(a animal) {
	fmt.Println("this animal: ", a.says())
}

type animal interface {
	says() string
	numberOfLegs() int
}

type dog struct {
	Name  string
	Breed string
}
type gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (d *dog) says() string {
	return "Dog!"
}

func (d *dog) numberOfLegs() int {
	return 4
}

func (g *gorilla) says() string {
	return "Gorilla!"
}
func (g *gorilla) numberOfLegs() int {
	return 2
}
