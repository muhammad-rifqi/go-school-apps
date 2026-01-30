package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"go-fiber-gorm/config"
	"go-fiber-gorm/models"
	"go-fiber-gorm/routes"
)

func main() {
	// Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	// Fiber
	app := fiber.New()

	// Routes
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
