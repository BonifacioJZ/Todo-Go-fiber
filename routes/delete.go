package routes

import (
	"sync"

	"github.com/BonifacioJZ/todo/db"
	"github.com/BonifacioJZ/todo/models"
	"github.com/gofiber/fiber/v2"
)

//funcion para eliminar una tarea.
func DeleteTask(c *fiber.Ctx) error {
	var id = c.Params("id")
	var task models.Task
	var db = db.ConnectionDB()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		db.Delete(&task, id)

	}()
	wg.Wait()

	return c.JSON("Eliminado")
}
