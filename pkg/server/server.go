package server

import (
	"net/http"
	"os"

	"github.com/nikndip/go_final_project/pkg/api"
)

func Run() error {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	mux := http.NewServeMux()
	api.Init(mux)
	mux.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(":"+port, mux)
}
