package middleware

import (
	"context"
	"jwt/utils"
	"net/http"
	"strings"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is blank", http.StatusBadRequest)
			return
		}

		// Extract the token from the "Bearer <token>" format
		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			http.Error(w, "Invalid Authorization header format", http.StatusBadRequest)
			return
		}
		token := tokenParts[1]

		id, err := utils.VerifyJwt(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Set the user ID in context instead of the header
		ctx := context.WithValue(r.Context(), "id", id)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
