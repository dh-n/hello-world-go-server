package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dh-n/go-course/pkg/config"
	"github.com/dh-n/go-course/pkg/handlers"
	"github.com/dh-n/go-course/pkg/render"
)

const portNumber = ":3000"

func main() {
	var app config.AppConfig

	tc,err := render.CacheTemplate()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting server on port, ", portNumber)
	// http.ListenAndServe(portNumber, nil)

	srv := http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

