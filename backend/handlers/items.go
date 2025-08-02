package handlers

import (
	"net/http"

	"ecommerce-backend/database"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
)

// CreateItem handles item creation
func CreateItem(c *gin.Context) {
	var req models.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default status if not provided
	if req.Status == "" {
		req.Status = "available"
	}

	item := models.Item{
		Name:   req.Name,
		Status: req.Status,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Item created successfully",
		"item":    item,
	})
}

// ListItems returns all items
func ListItems(c *gin.Context) {
	var items []models.Item
	if err := database.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
} 