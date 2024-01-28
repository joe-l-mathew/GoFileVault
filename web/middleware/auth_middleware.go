package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/joe-l-mathew/GoFileVault/utils"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/signup" || r.URL.Path == "/signin" {
			next.ServeHTTP(w, r)
			return
		}
		tokenString := extractTokenFromRequest(r)
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := utils.VerifyToken(tokenString)
		fmt.Println("clames", claims["userId"])
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check if the user exists in the database or perform any additional checks
		// For this example, assume the User model has a field 'Username'

		// Set the authenticated user in the request context
		ctx := context.WithValue(r.Context(), "userId", claims["userId"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	authParts := strings.Fields(authHeader)
	if len(authParts) != 1 {
		return ""
	}

	return authParts[0]
}
