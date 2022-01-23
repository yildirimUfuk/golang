package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type myIf interface {
	firstFunc()
	secondFunc()
}

type obj1 struct {
	name    string
	surname string
}

func (x *obj1) firstFunc() {
	fmt.Println("first func for obj1")
}
func (x *obj1) secondFunc() {
	fmt.Println("second func for obj2")
}

type obj2 struct {
	city string
}

func (x *obj2) firstFunc() {
	fmt.Println("first func for obj2")
}
func (x *obj2) secondFunc() {
	fmt.Println("second func for obj2")
}

func globalFunc(x myIf) {
	x.firstFunc()
}
func main() {
	// x := obj1{"ufuk", "yildirim"}
	y := obj2{"city"}
	globalFunc(&y)
}

//pointers example begin
func pointerEx() {
	type String *string

	str := "this is the string"
	var p String = &str
	*p = "asdf"

	fmt.Println(str, " - ", *p)
} //pointers example end

//read data from console example begin
func passOrFail() (string, error) {
	r := bufio.NewReader(os.Stdin)
	v, err := r.ReadString('\n')
	v = strings.TrimSpace(v)
	grade, err2 := strconv.Atoi(v)
	if err2 != nil {
		return v, errors.New("converting error")
	}

	if err != nil {
		return v, errors.New("reading error")
	}
	if grade < 50 {
		return "kaldi", nil
	}
	return "gecti", nil
} //read data from console example end
