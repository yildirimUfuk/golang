package httpPackage

import (
	"net/http"
)

var portNumber = ":8080"

// home is the home page handler
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.go.html")
}

//about is the about page handler
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.go.html")
}
