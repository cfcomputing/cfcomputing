package auth

import (
	"log"
	"net/http"
)

// again make sure to give a hat tip to https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81

// CheckAuth will check if the user is authenticated for the request based
func CheckAuth(reqContext string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Method: %v", r)

			authToken := r.Header.Get("Authorization")
			if len(authToken) <= 0 {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			log.Printf("authToken: %s", authToken)
			next.ServeHTTP(w, r)
		})

	}
}
