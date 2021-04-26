package routes

import (
	"sync"

	"github.com/BonifacioJZ/todo/db"
	"github.com/BonifacioJZ/todo/models"
	"github.com/gofiber/fiber/v2"
)

//Funcion para crear una tarea

func NewTask(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		db.ConnectionDB().Create(&models.Task{Name: task.Name, Description: task.Description, Done: task.Done})
	}()
	wg.Wait()

	return c.JSON(task)

}
