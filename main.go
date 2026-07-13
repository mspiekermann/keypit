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

	mux.HandleFunc("PUT /add", app.handleAddKeyValuePair)
	mux.HandleFunc("GET /all", app.handleGetAll)
	mux.HandleFunc("GET /get/{key}", app.handleGetByKey)
	mux.HandleFunc("GET /count", app.handleCountItems)

	return mux
}

func (app *App) handleAddKeyValuePair(w http.ResponseWriter, r *http.Request) {

	type AddKeyValuePairRequest struct {
		Key   string `json:"key"`
		Value any    `json:"value"`
	}

	var data AddKeyValuePairRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Printf("decoding request body: %v", err)
		http.Error(w, "invalid JSON request body", http.StatusBadRequest)
		return
	}

	app.store.Add(data.Key, data.Value)

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

func (app *App) handleCountItems(w http.ResponseWriter, r *http.Request) {

	data := map[string]int{
		"count": app.store.CountItems(),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("encoding response: %v", err)
	}
}
