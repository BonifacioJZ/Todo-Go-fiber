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

	app.Get("/api/tasks", routes.Allusers)
	app.Post("/api/newtask", routes.NewTask)

	app.Listen(":8000")

}
