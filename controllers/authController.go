package controllers

import "github.com/gofiber/fiber"

func Hello(c *fiber.Ctx) {
	c.SendString("Hello world !")
}

func Post(c *fiber.Ctx) {
	c.Status(200).JSON(&fiber.Map{
		"message": "Hello!",
	})
}
