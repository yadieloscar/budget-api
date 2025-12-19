// Package handlers contains HTTP handlers for the budget API.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CategoriesHandler handles category-related HTTP requests.
type CategoriesHandler struct{}

// NewCategoriesHandler returns a new categories handler.
func NewCategoriesHandler() *CategoriesHandler {
	return &CategoriesHandler{}
}

// GetCategories handles GET requests for retrieving categories
func (h *CategoriesHandler) GetCategories(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusOK, gin.H{"message": "GetCategories called"})
}

// CreateCategory handles POST requests for creating a new category
func (h *CategoriesHandler) CreateCategory(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusCreated, gin.H{"message": "CreateCategory called"})
}

// UpdateCategory handles PUT requests for updating an existing category
func (h *CategoriesHandler) UpdateCategory(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusOK, gin.H{"message": "UpdateCategory called"})
}

// DeleteCategory handles DELETE requests for removing a category
func (h *CategoriesHandler) DeleteCategory(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusNoContent, nil)
}

// RegisterCategoriesRoutes mounts the categories endpoints under the provided router group.
func RegisterCategoriesRoutes(r *gin.RouterGroup, h *CategoriesHandler) {
	api := r.Group("/categories")
	{
		api.GET("/", h.GetCategories)
		api.POST("/", h.CreateCategory)
		api.PUT("/:id", h.UpdateCategory)
		api.DELETE("/:id", h.DeleteCategory)
	}
}
