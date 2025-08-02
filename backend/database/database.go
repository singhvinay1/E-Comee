package database

import (
	"log"

	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// InitDatabase initializes the database connection and migrates the schema
func InitDatabase() {
	var err error
	
	// Use configuration for database connection
	dbType := config.AppConfig.Database.Type
	dbName := config.AppConfig.Database.Name
	
	DB, err = gorm.Open(dbType, dbName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Enable GORM logging
	DB.LogMode(true)

	// Auto migrate the schema
	err = DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
	).Error

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Seed initial data
	seedInitialData()
}

// seedInitialData adds some initial items to the database
func seedInitialData() {
	var count int
	DB.Model(&models.Item{}).Count(&count)
	
	if count == 0 {
		items := []models.Item{
			{Name: "Laptop", Status: "available"},
			{Name: "Smartphone", Status: "available"},
			{Name: "Headphones", Status: "available"},
			{Name: "Tablet", Status: "available"},
			{Name: "Wireless Mouse", Status: "available"},
			{Name: "Keyboard", Status: "available"},
			{Name: "Monitor", Status: "available"},
			{Name: "USB Cable", Status: "available"},
		}

		for _, item := range items {
			DB.Create(&item)
		}
		log.Println("Initial items seeded successfully")
	}
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
} 