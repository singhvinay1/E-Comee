package routes

import (
	"ecommerce-backend/config"
	"ecommerce-backend/handlers"
	"ecommerce-backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", config.AppConfig.CORS.Origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// User routes
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.ListUsers)
	r.POST("/users/login", handlers.Login)

	// Item routes
	r.POST("/items", handlers.CreateItem)
	r.GET("/items", handlers.ListItems)

	// Cart routes (protected)
	cartRoutes := r.Group("/carts")
	cartRoutes.Use(middleware.AuthMiddleware())
	{
		cartRoutes.POST("/", handlers.AddToCart)
		cartRoutes.GET("/my", handlers.GetUserCart)
		cartRoutes.DELETE("/clear", handlers.ClearCart)
		cartRoutes.DELETE("/remove", handlers.RemoveFromCart)
	}
	r.GET("/carts", handlers.ListCarts) // Public endpoint

	// Order routes (protected)
	orderRoutes := r.Group("/orders")
	orderRoutes.Use(middleware.AuthMiddleware())
	{
		orderRoutes.POST("/", handlers.CreateOrder)
		orderRoutes.GET("/my", handlers.GetUserOrders)
	}
	r.GET("/orders", handlers.ListOrders) // Public endpoint

	return r
} 