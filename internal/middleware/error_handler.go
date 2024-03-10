package middleware

import (
	"com.amagine.router/internal/router"
	"net/http"
)

// ErrorHandler is a middleware function for handling errors
func ErrorHandler() router.Middleware {
	return func(handler router.HandlerFunc) router.HandlerFunc {
		return func(c router.Context) (interface{}, error) {
			result, err := handler(c)

			// If there's an error, handle it
			if err != nil {
				// Log the error
				c.JSON(http.StatusInternalServerError, router.Response{
					Error: err.Error(),
				})

				return nil, err
			}

			return result, nil
		}
	}
}
