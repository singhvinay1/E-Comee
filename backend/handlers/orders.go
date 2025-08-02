package handlers

import (
	"net/http"

	"ecommerce-backend/database"
	"ecommerce-backend/middleware"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
)

// CreateOrder handles converting a cart to an order
func CreateOrder(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify the cart belongs to the user
	var cart models.Cart
	if err := database.DB.Where("id = ? AND user_id = ?", req.CartID, user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found or doesn't belong to user"})
		return
	}

	// Check if cart is active
	if cart.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is not active"})
		return
	}

	// Check if cart has items
	var cartItems []models.CartItem
	if err := database.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	// Create order
	order := models.Order{
		CartID: cart.ID,
		UserID: user.ID,
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Remove all items from the cart
	if err := database.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart items"})
		return
	}

	// Update cart status to "converted"
	cart.Status = "converted"
	database.DB.Save(&cart)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

// ListOrders returns all orders
func ListOrders(c *gin.Context) {
	var orders []models.Order
	if err := database.DB.Preload("User").Preload("Cart").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetUserOrders returns orders for the current user
func GetUserOrders(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	var orders []models.Order
	if err := database.DB.Where("user_id = ?", user.ID).Preload("Cart").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user orders"})
		return
	}

	// Load cart items for each order
	for i := range orders {
		if orders[i].Cart != nil {
			var cartItems []models.CartItem
			if err := database.DB.Where("cart_id = ?", orders[i].Cart.ID).Preload("Item").Find(&cartItems).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
				return
			}

			// Extract items from cart items
			var items []models.Item
			for _, cartItem := range cartItems {
				if cartItem.Item != nil {
					items = append(items, *cartItem.Item)
				}
			}
			orders[i].Cart.Items = items
		}
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
} 