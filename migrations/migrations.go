package migrations

import (
	"github.com/BonifacioJZ/todo/db"
	"github.com/BonifacioJZ/todo/models"
)

func Initmigrations() {
	db.ConnectionDB().AutoMigrate(&models.Task{})
}
