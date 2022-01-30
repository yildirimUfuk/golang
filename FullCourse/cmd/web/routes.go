package web

import (
	"golangTutorial/pkg/config"
	"golangTutorial/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// //this routes funciton for bmizerany package routing
// func routes(app *config.AppConfig) http.Handler {
// 	mux := pat.New()
// 	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
// 	return mux
// }

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	// mux.Use(WriteToConsole)

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
