package main

import (
	"fmt"
	"net/http"
)

func (app *App) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Keyp it!")
	})

	mux.HandleFunc("PUT /items/{key}", app.handleSetKeyValuePair)
	mux.HandleFunc("GET /items", app.handleGetAll)
	mux.HandleFunc("GET /items/{key}", app.handleGetByKey)
	mux.HandleFunc("DELETE /items/{key}", app.handleDeleteKeyValuePair)
	mux.HandleFunc("GET /stats", app.handleStats)

	return mux
}
