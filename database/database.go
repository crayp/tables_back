package database

import (
	"fmt"
	"log"

	"restaurant-booking-backend/config"
	"restaurant-booking-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	// Автоматическая миграция моделей
	DB.AutoMigrate(&models.User{}, &models.Restaurant{}, &models.Table{}, &models.Booking{})
}
