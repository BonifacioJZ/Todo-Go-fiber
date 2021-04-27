package routes

import (
	"sync"

	"github.com/BonifacioJZ/todo/db"
	"github.com/BonifacioJZ/todo/models"
	"github.com/gofiber/fiber/v2"
)

// funcion para canmbiar el estado de done
func Done(c *fiber.Ctx) error {
	var id = c.Params("id")
	var task models.Task
	var newData models.Task
	var wg sync.WaitGroup
	if err := c.BodyParser(&newData); err != nil {
		return err
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		db.ConnectionDB().First(&task, id)
		task.Done = newData.Done
		db.ConnectionDB().Save(&task)
	}()
	wg.Wait()

	return c.JSON(task)
}
