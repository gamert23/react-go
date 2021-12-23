package main

import (
	"crud/database"

	"crud/routes"

	"github.com/gofiber/fiber"
)

func main() {
	// Start Connect DB
	database.Connect()
	// End Connect DB

	app := fiber.New()

	// Set Routes
	routes.Setup(app)
	// End Set Routes

	app.Listen(8081)
}
