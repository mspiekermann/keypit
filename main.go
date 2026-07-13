package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Keyp it!")
	})

	server := &http.Server{
		Handler: mux,
		Addr:    ":9000",
	}

	server.ListenAndServe()

}
