package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
)

func CheckPermissions(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			GenerateErrorResponse(w, "Não authorizado", "Token é necessario", http.StatusUnauthorized)
			return
		}
		token = strings.Split(token, " ")[1]
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			GenerateErrorResponse(w, "Não authorizado", "Token não é valido", http.StatusUnauthorized)
			return
		}
		client := &http.Client{}
		urlAuth := os.Getenv("URL_AUTH")
		request, err := http.NewRequest(http.MethodGet, urlAuth, nil)
		request.Header.Set("Authorization", "Bearer "+token)
		client.Timeout = time.Second * 10
		response, err := client.Do(request)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			GenerateErrorResponse(w, "Auth Error", err.Error(), http.StatusForbidden0)
			return
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
