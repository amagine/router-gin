package main

import (
	"com.amagine.router/internal/handlers"
	"com.amagine.router/internal/middleware"
	"com.amagine.router/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new instance of http.Engine
	engine := gin.Default()

	// common middleware
	middlewares := make([]router.Middleware, 0)
	middlewares = append(middlewares, middleware.Logger())
	middlewares = append(middlewares, middleware.ErrorHandler())

	// Create a new instance of RouteGroup
	routerGroup := router.NewRouteGroup("/v1", engine, middlewares...)

	// Define router
	routerGroup.AddHandler("GET", "/ping", handlers.PingHandler)
	routerGroup.AddHandler("GET", "/user/:id", handlers.GetUserHandler, middleware.Logger(), middleware.Authentication())

	// Start the server
	engine.Run(":8080")
}
