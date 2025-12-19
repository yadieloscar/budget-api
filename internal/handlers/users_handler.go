// Package handlers contains HTTP handlers for the budget API.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UsersHandler handles user-related HTTP requests.
type UsersHandler struct{}

// NewUsersHandler returns a new users handler.
func NewUsersHandler() *UsersHandler {
	return &UsersHandler{}
}

// GetUser handles GET requests for retrieving user information
func (h *UsersHandler) GetUser(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusOK, gin.H{"message": "GetUser called"})
}

// CreateUser handles POST requests for creating a new user
func (h *UsersHandler) CreateUser(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusCreated, gin.H{"message": "CreateUser called"})
}

// UpdateUser handles PUT requests for updating an existing user
func (h *UsersHandler) UpdateUser(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusOK, gin.H{"message": "UpdateUser called"})
}

// DeleteUser handles DELETE requests for removing a user
func (h *UsersHandler) DeleteUser(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusNoContent, nil)
}

// RegisterUsersRoutes mounts the users endpoints under the provided router group.
func RegisterUsersRoutes(r *gin.RouterGroup, h *UsersHandler) {
	api := r.Group("/users")
	{
		api.GET("/:id", h.GetUser)
		api.POST("/", h.CreateUser)
		api.PUT("/:id", h.UpdateUser)
		api.DELETE("/:id", h.DeleteUser)
	}
}
