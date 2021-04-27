package routes

import (
	"sync"

	"github.com/BonifacioJZ/todo/db"
	"github.com/BonifacioJZ/todo/models"
	"github.com/gofiber/fiber/v2"
)

//Funcion pra monstrar una tarea
func Task(c *fiber.Ctx) error {
	var id = c.Params("id")
	var wg sync.WaitGroup
	var task models.Task
	wg.Add(1)
	go func() {
		defer wg.Done()
		db.ConnectionDB().First(&task, id)
	}()
	wg.Wait()

	return c.JSON(task)
}
