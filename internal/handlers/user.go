package handlers

import "com.amagine.router/internal/router"

// UserResponse represents the structure of the user response
type UserResponse struct {
	UserID string `json:"userID"`
}

// GetUserHandler handles the /user/:id route
func GetUserHandler(c router.Context) (interface{}, error) {
	userID := c.Param("id")
	return UserResponse{
		UserID: userID,
	}, nil
}
