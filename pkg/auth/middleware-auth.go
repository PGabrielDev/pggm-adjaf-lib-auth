package auth

import (
	"encoding/json"
	"net/http"
	"strings"
)

func CheckPermissions(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			GenerateErrorResponse(w, "Não authorizado", "Token é necessario", http.StatusUnauthorized)
		}
		token = strings.Split(token, " ")[1]
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			GenerateErrorResponse(w, "Não authorizado", "Token não é valido", http.StatusUnauthorized)
		}

		next(w, r)
	}
}

func GenerateErrorResponse(w http.ResponseWriter, message, description string, statusCode int) {
	json.NewEncoder(w).Encode(struct {
		Message     string `json:"message"`
		Description string `json:"description"`
		Status      int    `json:"status"`
	}{
		Message:     message,
		Description: description,
		Status:      statusCode,
	})
}
