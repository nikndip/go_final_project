package api

import "net/http"

func Init(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/nextdate", nextDateHandler)
	mux.HandleFunc("POST /api/task", addTaskHandler)
}
