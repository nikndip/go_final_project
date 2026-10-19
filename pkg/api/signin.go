package api

import (
	"crypto/hmac"
	"encoding/json"
	"net/http"
	"os"
)

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Некорректный JSON",
		})
		return
	}

	password := os.Getenv("TODO_PASSWORD")
	if password == "" || !hmac.Equal(
		[]byte(request.Password),
		[]byte(password),
	) {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "Неверный пароль",
		})
		return
	}

	token, err := createToken(password)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
