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

	app.Get("/api/task", routes.Allusers)
	app.Get("/api/task/:id", routes.Task)
	app.Post("/api/task", routes.NewTask)
	app.Put("/api/task/:id", routes.UpdateTasks)
	app.Put("/api/done/:id", routes.Done)
	app.Delete("/api/task/:id", routes.DeleteTask)

	app.Listen(":8000")

}
