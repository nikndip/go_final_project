package api

import "net/http"

func Init(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/nextdate", nextDateHandler)
	mux.HandleFunc("POST /api/task", addTaskHandler)
	mux.HandleFunc("GET /api/tasks", tasksHandler)
	mux.HandleFunc("GET /api/task", getTaskHandler)
	mux.HandleFunc("PUT /api/task", editTaskHandler)
	mux.HandleFunc("POST /api/task/done", doneTaskHandler)
	mux.HandleFunc("DELETE /api/task", deleteTaskHandler)
}
