package main

import (
	"encoding/json"
	"net/http"
)

type setValueRequest struct {
	Value any `json:"value"`
}

func (app *App) handleSetKeyValuePair(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	var request setValueRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		app.logger.Printf("decoding request body: %v", err)
		http.Error(w, "invalid JSON request body", http.StatusBadRequest)
		return
	}

	app.store.Add(key, request.Value)
	w.WriteHeader(http.StatusCreated)
}

func (app *App) handleGetAll(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, app.store.GetAll())
}

func (app *App) handleGetByKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	value, exists := app.store.Get(key)
	if !exists {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}

	app.writeJSON(w, http.StatusOK, value)
}

func (app *App) handleStats(w http.ResponseWriter, r *http.Request) {
	response := map[string]int{
		"count": app.store.CountItems(),
	}

	app.writeJSON(w, http.StatusOK, response)
}

func (app *App) handleDeleteKeyValuePair(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	if !app.store.Delete(key) {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *App) writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		app.logger.Printf("encoding response: %v", err)
	}
}
