package middleware

import "net/http"

// Middleware represents a middleware
type Middleware func(http.Handler) http.Handler

// Adapt group the middlewares on request
func Adapt(next http.Handler, chains ...Middleware) http.Handler {
	for _, item := range chains {
		next = item(next)
	}
	return next
}
