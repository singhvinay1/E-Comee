package handlers

import (
	"net/http"

	"ecommerce-backend/database"
	"ecommerce-backend/middleware"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
)

// AddToCart handles adding items to a user's cart
func AddToCart(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	var req models.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get or create user's cart
	var cart models.Cart
	if err := database.DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
		// Create new cart if it doesn't exist
		cart = models.Cart{
			UserID: user.ID,
			Name:   "My Cart",
			Status: "active",
		}
		if err := database.DB.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
			return
		}
	}

	// Add items to cart
	for _, itemID := range req.ItemIDs {
		// Check if item exists
		var item models.Item
		if err := database.DB.First(&item, itemID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found"})
			return
		}

		// Check if item is already in cart
		var existingCartItem models.CartItem
		if err := database.DB.Where("cart_id = ? AND item_id = ?", cart.ID, itemID).First(&existingCartItem).Error; err == nil {
			// Item already in cart, skip
			continue
		}

		// Add item to cart
		cartItem := models.CartItem{
			CartID: cart.ID,
			ItemID: itemID,
		}
		if err := database.DB.Create(&cartItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
			return
		}
	}

	// Update user's cart ID
	user.CartID = &cart.ID
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Items added to cart successfully",
		"cart_id": cart.ID,
	})
}

// ListCarts returns all carts
func ListCarts(c *gin.Context) {
	var carts []models.Cart
	if err := database.DB.Preload("User").Preload("Items").Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch carts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"carts": carts})
}

// GetUserCart returns the current user's cart
func GetUserCart(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Load cart items with item details
	var cartItems []models.CartItem
	if err := database.DB.Where("cart_id = ?", cart.ID).Preload("Item").Find(&cartItems).Error; err != nil {
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

	cart.Items = items

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

// RemoveFromCart removes a specific item from the user's cart
func RemoveFromCart(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	var req struct {
		ItemID uint `json:"item_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Delete the specific cart item
	if err := database.DB.Where("cart_id = ? AND item_id = ?", cart.ID, req.ItemID).Delete(&models.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}

// ClearCart removes all items from the user's cart
func ClearCart(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Delete all cart items for this cart
	if err := database.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared successfully"})
} 