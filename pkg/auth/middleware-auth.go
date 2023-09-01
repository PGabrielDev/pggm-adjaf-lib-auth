package auth

import (
	"encoding/json"
	"github.com/PGabrielDev/pggm--adjaf-lib-auth/pkg/auth/DTOs"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func CheckPermissions(next http.HandlerFunc, dto DTOs.AuthPermission) http.HandlerFunc {
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
			GenerateErrorResponse(w, "Auth Error", err.Error(), http.StatusForbidden)
			return
		}
		defer response.Body.Close()
		responsePayload, err := io.ReadAll(response.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			GenerateErrorResponse(w, "Parser error", err.Error(), http.StatusInternalServerError)
			return
		}
		var authDTO []DTOs.UserAccessDTO
		if err := json.Unmarshal(responsePayload, &authDTO); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			GenerateErrorResponse(w, "Parser error", err.Error(), http.StatusInternalServerError)
			return
		}
		for _, v := range authDTO {
			if dto.Name == v.ProductName {
				for _, level := range v.LevelAccess {
					if dto.Permission.String() == level.LevelAccessName {
						next(w, r)
					}
				}
			}
		}
		w.WriteHeader(http.StatusForbidden)
		GenerateErrorResponse(w, "Auth error", "Sem permissoes necessarias", http.StatusForbidden)
		return
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

func containLevelAccess(level string, listAccess []DTOs.LevelAccessDTO) {

}
