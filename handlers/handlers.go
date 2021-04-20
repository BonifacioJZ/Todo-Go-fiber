package handlers

import (
	"github.com/BonifacioJZ/todo/routes"
	"github.com/gofiber/fiber/v2"
)

func Handlers() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Get("/tasks", routes.Allusers)

	app.Listen(":8000")

}
