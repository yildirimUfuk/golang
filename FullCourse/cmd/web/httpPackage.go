package web

import (
	"fmt"
	"golangTutorial/pkg/config"
	"golangTutorial/pkg/handlers"
	"golangTutorial/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

//main http package example function.
func HttpMain() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

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
