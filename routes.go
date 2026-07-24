package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func (app *App) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Keyp it!")
	})

	mux.HandleFunc("GET /health", app.handleHealthcheck)

	mux.HandleFunc("PUT /items/{key}", app.handleSetKeyValuePair)
	mux.HandleFunc("GET /items", app.handleGetAll)
	mux.HandleFunc("GET /items/{key}", app.handleGetByKey)
	mux.HandleFunc("DELETE /items/{key}", app.handleDeleteKeyValuePair)
	mux.HandleFunc("GET /stats", app.handleStats)
	mux.HandleFunc("POST /snapshot", app.handleSnapshotToFile)

	return requestLogger(mux)
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.DebugContext(
			r.Context(),
			"HTTP request received",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("user_agent", r.UserAgent()),
		)

		next.ServeHTTP(w, r)
	})
}
