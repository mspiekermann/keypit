package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	store  Store
	logger *log.Logger
}

func main() {

	app := App{
		store:  *NewStore(),
		logger: log.Default(),
	}

	mux := app.routes()

	server := &http.Server{
		Handler: mux,
		Addr:    ":9000",
	}

	server.ListenAndServe()
}

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

func (app *App) handleSetKeyValuePair(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	type SetValueRequest struct {
		Value any `json:"value"`
	}

	var data SetValueRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Printf("decoding request body: %v", err)
		http.Error(w, "invalid JSON request body", http.StatusBadRequest)
		return
	}

	app.store.Add(key, data.Value)

	w.WriteHeader(http.StatusCreated)
}

func (app *App) handleGetAll(w http.ResponseWriter, r *http.Request) {

	data := app.store.GetAll()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("encoding response: %v", err)
	}
}

func (app *App) handleGetByKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	value, exists := app.store.Get(key)
	if !exists {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("encoding response: %v", err)
	}
}

func (app *App) handleStats(w http.ResponseWriter, r *http.Request) {

	data := map[string]int{
		"count": app.store.CountItems(),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("encoding response: %v", err)
	}
}

func (app *App) handleDeleteKeyValuePair(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	if !app.store.Delete(key) {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
