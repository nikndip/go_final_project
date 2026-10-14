package server

import (
	"net/http"
	"os"
)

func Run() error {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(":"+port, mux)
}
