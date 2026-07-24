package main

import (
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	store *Store
}

func main() {

	loggerConfiguration()

	app := App{
		store: NewStore(),
	}

	server := &http.Server{
		Handler: app.routes(),
		Addr:    ":9000",
	}

	slog.Info(
		"server started",
		"port", server.Addr,
	)

	server.ListenAndServe()
}

func loggerConfiguration() {
	level := slog.LevelInfo

	if os.Getenv("APP_ENV") == "development" {
		level = slog.LevelDebug
	}

	slog.SetDefault(slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}),
	))
}
