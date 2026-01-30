package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/users", handlers.CreateUser)
	api.Get("/users", handlers.GetUsers)
}
