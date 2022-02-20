package main

import (
	"fmt"
	"golangTutorial/pkg/handlers"
	"net/http"
)

type indicies struct {
	x int
	y int
	z int
}

func main() {
	HttpMain()
}

var portNumber = ":8080"

//main http package example function.
func HttpMain() {
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))
	// fmt.Printf(fmt.Sprintf("starting application on port: %s\n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
