package router

import "github.com/gin-gonic/gin"

// ginContextAdapter adapts Gin's context to our custom Context interface
type ginContextAdapter struct {
	*gin.Context
}

// Param retrieves the value of the URL parameter
func (c ginContextAdapter) Param(key string) string {
	return c.Context.Param(key)
}

// JSON marshals the given struct as JSON into the response body
func (c ginContextAdapter) JSON(code int, obj interface{}) {
	c.Context.JSON(code, obj)
}
