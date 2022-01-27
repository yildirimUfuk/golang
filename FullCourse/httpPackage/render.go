package httpPackage

import (
	"fmt"
	"net/http"
	"text/template"
)

//this render function gets a http.ResponseWriter and the page name and it renders the page.
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
