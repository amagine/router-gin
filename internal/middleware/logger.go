package middleware

import "com.amagine.router/internal/router"

// Logger is a middleware function for logging
func Logger() router.Middleware {
	return func(handler router.HandlerFunc) router.HandlerFunc {
		return func(c router.Context) (interface{}, error) {
			// Do something before handling the request
			result, err := handler(c)
			// Do something after handling the request
			return result, err
		}
	}
}
