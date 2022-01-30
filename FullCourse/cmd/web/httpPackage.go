package web

import (
	"fmt"
	"golangTutorial/pkg/config"
	"golangTutorial/pkg/handlers"
	"golangTutorial/pkg/render"
	"log"
	"net/http"
)

var portNumber = ":8080"

//main http package example function.
func HttpMain() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot crate template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))
	// fmt.Printf(fmt.Sprintf("starting application on port: %s\n", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("server connot opened! ", err)
	}
}
