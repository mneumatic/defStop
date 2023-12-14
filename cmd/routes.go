package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/mneumatic/defstop/models"
	config "github.com/mneumatic/defstop/pkg/configs"
	"github.com/mneumatic/defstop/pkg/handlers"
	"github.com/mneumatic/defstop/pkg/render"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//  CHI NotFound & MethodNotAllowed ERROR HANDLING
	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		stringMap := make(map[string]string)
		stringMap["Title"] = "Page Not Found | defStop"
		render.RenderTemplate(w, r, "error.tmpl", &models.TemplateData{
			StringMap: stringMap,
		})
		w.WriteHeader(404)
	})

	mux.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		stringMap := make(map[string]string)
		stringMap["Title"] = "Page Not Found | defStop"
		render.RenderTemplate(w, r, "error.tmpl", &models.TemplateData{
			StringMap: stringMap,
		})
		w.WriteHeader(405)
	})

	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return mux
}
