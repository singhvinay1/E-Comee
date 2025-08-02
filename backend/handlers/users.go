package handlers

import (
	"net/http"

	"ecommerce-backend/database"
	"ecommerce-backend/models"
	"ecommerce-backend/utils"

	"github.com/gin-gonic/gin"
)

// CreateUser handles user registration
func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username already exists
	var existingUser models.User
	if err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Generate token
	token, err := utils.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Create user
	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Token:    token,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Create cart for the user
	cart := models.Cart{
		UserID: user.ID,
		Name:   "My Cart",
		Status: "active",
	}

	if err := database.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}

	// Update user with cart ID
	user.CartID = &cart.ID
	database.DB.Save(&user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

// Login handles user authentication
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by username
	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate new token (single device login)
	newToken, err := utils.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Update user token
	user.Token = newToken
	database.DB.Save(&user)

	c.JSON(http.StatusOK, models.LoginResponse{
		Token: newToken,
		User:  user,
	})
}

// ListUsers returns all users
func ListUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
} 