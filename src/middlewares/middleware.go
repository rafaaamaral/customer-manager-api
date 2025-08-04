package middlewares

import (
	"customer-manager-api/src/authentication"
	"fmt"
	"net/http"
)

func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		nextFunction(w, r)
	}
}

func Authorize(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authorizing request...")
		if err := authentication.ValidateToken(r); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		nextFunction(w, r)
	}
}
