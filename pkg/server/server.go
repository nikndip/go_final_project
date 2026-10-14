package server

import "net/http"

func Run() error {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(":7540", mux)
}
