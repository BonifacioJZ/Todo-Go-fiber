package routes

import (
	"sync"

	"github.com/BonifacioJZ/todo/db"
	"github.com/BonifacioJZ/todo/models"
	"github.com/gofiber/fiber/v2"
)

func Allusers(c *fiber.Ctx) error {

	var tasks []models.Task

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		db.ConnectionDB().Find(&tasks)

	}()
	wg.Wait()

	return c.JSON(tasks)

}
