package main

import (
	"github.com/ermos/annotation"
	"github.com/ermos/annotation/parser"
	"github.com/vitego/boilerplate"
	"github.com/vitego/boilerplate/internal/app/exceptions"
	"github.com/vitego/boilerplate/internal/app/http/controllers"
	"github.com/vitego/boilerplate/internal/app/http/middlewares"
	"github.com/vitego/config"
	"github.com/vitego/router"
	"log"
)

func main() {
	var routes []parser.API

	// Parse config files
	err := config.Parse("config", boilerplate.Config)
	if err != nil {
		log.Fatal(err)
	}

	// Generate annotation
	if config.Get("app.debug") == "true" {
		err = annotation.Fetch("internal/app/http/controllers", &routes, parser.ToAPI)
		if err != nil {
			log.Fatal(err)
		}

		err = annotation.Save(routes, ".build/routes.json")
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Fatal(router.ListenAndServe(
		boilerplate.Routes,
		controllers.Handler{},
		exceptions.Handler{},
		middlewares.Handler{},
	))
}
