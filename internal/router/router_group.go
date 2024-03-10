package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Context represents the interface for the HTTP request context
type Context interface {
	Param(key string) string
	JSON(code int, obj interface{})
}

// HandlerFunc represents the interface for handler functions
type HandlerFunc func(c Context) (interface{}, error)

// Response represents the structure of the response
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// Middleware represents the interface for middleware functions
type Middleware func(handler HandlerFunc) HandlerFunc

// RouteGroup represents a group of router
type RouteGroup struct {
	group      *gin.RouterGroup
	middleware []Middleware
}

// NewRouteGroup creates a new instance of RouteGroup
func NewRouteGroup(prefix string, engine *gin.Engine, globalMiddleware ...Middleware) *RouteGroup {
	group := engine.Group(prefix)
	return &RouteGroup{
		group:      group,
		middleware: globalMiddleware,
	}
}

// Use adds middleware to the route group
func (r *RouteGroup) Use(middleware ...Middleware) {
	r.middleware = append(r.middleware, middleware...)
}

// AddHandler adds a new route and handler to the route group
func (r *RouteGroup) AddHandler(method, path string, handler HandlerFunc, middleware ...Middleware) {
	combinedMiddleware := append(r.middleware, middleware...)
	r.group.Handle(method, path, func(c *gin.Context) {
		contextAdapter := ginContextAdapter{
			Context: c,
		}
		handlerFunc := composeMiddleware(handler, combinedMiddleware)
		data, err := handlerFunc(contextAdapter)
		if err != nil {
			return // Errorhandler middleware is responsible for error handling
		}
		contextAdapter.JSON(http.StatusOK, Response{
			Data: data,
		})
	})
}

// composeMiddleware composes multiple middleware functions into a single handler function
func composeMiddleware(handler HandlerFunc, middleware []Middleware) HandlerFunc {
	for _, mw := range middleware {
		handler = mw(handler)
	}
	return handler
}
