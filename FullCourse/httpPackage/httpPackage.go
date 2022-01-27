package httpPackage

import (
	"fmt"
	"net/http"
)

//main http package example function.
func HttpMain() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))
	// fmt.Printf(fmt.Sprintf("starting application on port: %s\n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
