package handlers

import "com.amagine.router/internal/router"

// PingResponse represents the structure of the ping response
type PingResponse struct {
	Message string `json:"message"`
}

// PingHandler handles the /ping route
func PingHandler(c router.Context) (interface{}, error) {
	return PingResponse{
		Message: "pong",
	}, nil
}
