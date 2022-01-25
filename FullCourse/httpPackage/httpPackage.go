package httppackage

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"

// home is the home page handler
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home func")
}

//about is the about page handler
func about(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "this is the about page and 2+2=%d", addVal(2, 2))
}

//sum two integer and returns the sum of them
func addVal(x, y int) int {
	return x + y
}

//main http package example function.
func HttpMain() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
