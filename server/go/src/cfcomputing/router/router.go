package router

import (
	"net/http"
)

// Hat tip to: https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81

// Middleware type used to wrap our http handler
type Middleware func(http.Handler) http.Handler

// ApplyMiddleware will execute the middleware in the order they were passed in
func ApplyMiddleware(h http.Handler, middleware ...Middleware) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}
