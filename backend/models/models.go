package models

import (
	"time"
)

// User represents a user account
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // "-" means this field won't be included in JSON
	Token     string    `json:"token" gorm:"unique"`
	CartID    *uint     `json:"cart_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relationships
	Cart   *Cart    `json:"cart,omitempty" gorm:"foreignkey:CartID"`
	Orders []Order  `json:"orders,omitempty" gorm:"foreignkey:UserID"`
}

// Item represents a product/item in the store
type Item struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"not null"`
	Status    string    `json:"status" gorm:"default:'available'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relationships
	CartItems []CartItem `json:"cart_items,omitempty" gorm:"foreignkey:ItemID"`
}

// Cart represents a user's shopping cart
type Cart struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id" gorm:"not null;unique"`
	Name      string    `json:"name"`
	Status    string    `json:"status" gorm:"default:'active'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relationships
	User      *User       `json:"user,omitempty" gorm:"foreignkey:UserID"`
	Items     []Item      `json:"items,omitempty" gorm:"many2many:cart_items;"`
	CartItems []CartItem  `json:"cart_items,omitempty" gorm:"foreignkey:CartID"`
	Order     *Order      `json:"order,omitempty" gorm:"foreignkey:CartID"`
}

// CartItem represents the junction table between Cart and Item
type CartItem struct {
	CartID    uint      `json:"cart_id" gorm:"primary_key"`
	ItemID    uint      `json:"item_id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	
	// Relationships
	Cart *Cart `json:"cart,omitempty" gorm:"foreignkey:CartID"`
	Item *Item `json:"item,omitempty" gorm:"foreignkey:ItemID"`
}

// Order represents a placed order
type Order struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CartID    uint      `json:"cart_id" gorm:"not null;unique"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relationships
	Cart *Cart `json:"cart,omitempty" gorm:"foreignkey:CartID"`
	User *User `json:"user,omitempty" gorm:"foreignkey:UserID"`
}

// Request/Response structures

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// CreateUserRequest represents the user creation request
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateItemRequest represents the item creation request
type CreateItemRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status"`
}

// AddToCartRequest represents the add to cart request
type AddToCartRequest struct {
	ItemIDs []uint `json:"item_ids" binding:"required"`
}

// CreateOrderRequest represents the order creation request
type CreateOrderRequest struct {
	CartID uint `json:"cart_id" binding:"required"`
} 