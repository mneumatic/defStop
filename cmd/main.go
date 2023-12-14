package main

import (
	"fmt"
	config "github.com/mneumatic/defstop/pkg/configs"
	"github.com/mneumatic/defstop/pkg/handlers"
	"github.com/mneumatic/defstop/pkg/render"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var app config.AppConfig

	// Production mode
	app.Production = false

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache.")
	} else {
		log.Println("Template cache created.")
	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// Check if PORT is available.
	// If "false" set port
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	// Format port string
	addr := fmt.Sprintf(":%s", port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      routes(&app),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("main: running server on port", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("main: couldn't start server: %v\n", err)
	}
}
