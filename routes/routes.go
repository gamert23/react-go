package routes

import (
	"crud/controllers"

	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	// Test get
	app.Get("/", controllers.Hello)

	// Test post
	app.Post("/post", controllers.Post)
}
