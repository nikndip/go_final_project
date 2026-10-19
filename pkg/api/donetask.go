package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/nikndip/go_final_project/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		status := http.StatusInternalServerError
		message := err.Error()

		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusNotFound
			message = "Задача не найдена"
		}

		writeJSON(w, status, map[string]string{"error": message})
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
	} else {
		next, dateErr := NextDate(time.Now(), task.Date, task.Repeat)
		if dateErr != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{
				"error": dateErr.Error(),
			})
			return
		}

		err = db.UpdateDate(next, id)
	}

	if err != nil {
		status := http.StatusInternalServerError
		message := err.Error()

		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusNotFound
			message = "Задача не найдена"
		}

		writeJSON(w, status, map[string]string{"error": message})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{})
}
