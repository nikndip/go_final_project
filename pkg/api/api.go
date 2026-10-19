package api

import "net/http"

func Init(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/signin", signinHandler)
	mux.HandleFunc("GET /api/nextdate", nextDateHandler)

	mux.HandleFunc("POST /api/task", auth(addTaskHandler))
	mux.HandleFunc("GET /api/task", auth(getTaskHandler))
	mux.HandleFunc("PUT /api/task", auth(editTaskHandler))
	mux.HandleFunc("DELETE /api/task", auth(deleteTaskHandler))
	mux.HandleFunc("GET /api/tasks", auth(tasksHandler))
	mux.HandleFunc("POST /api/task/done", auth(doneTaskHandler))
}
