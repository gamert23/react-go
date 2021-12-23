package controllers

import "github.com/gofiber/fiber"

// Hello via '/'
func Hello(c *fiber.Ctx) {
	c.SendString("Hello world !")
}

// Post via '/post'
func Post(c *fiber.Ctx) {
	c.Status(200).JSON(&fiber.Map{
		"message": "Hello!",
	})
}
