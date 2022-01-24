package main

import (
	"bufio"
	"errors"
	"fmt"
	"golangTutorial/myPackage"
	"os"
	"strconv"
	"strings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(myPackage.MyPackageFunc())
	fmt.Println(quote.Go())
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
