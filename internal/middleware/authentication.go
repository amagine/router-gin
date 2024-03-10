package middleware

import "com.amagine.router/internal/router"

// Authentication is a middleware function for authentication
func Authentication() router.Middleware {
	return func(handler router.HandlerFunc) router.HandlerFunc {
		return func(c router.Context) (interface{}, error) {
			// Check authentication logic here
			// For demonstration purposes, let's assume authentication is successful
			return handler(c)
		}
	}
}
