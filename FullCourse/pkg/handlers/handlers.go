package handlers

import (
	"golangTutorial/pkg/render"
	"net/http"
)

// home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.go.html")
}

//about is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.go.html")
}
