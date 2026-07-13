package main

import (
	"log"
	"net/http"
)

type App struct {
	store  *Store
	logger *log.Logger
}

func main() {

	app := App{
		store:  NewStore(),
		logger: log.Default(),
	}

	server := &http.Server{
		Handler: app.routes(),
		Addr:    ":9000",
	}

	server.ListenAndServe()
}
