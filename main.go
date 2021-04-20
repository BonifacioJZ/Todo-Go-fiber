package main

import (
	"log"

	"github.com/BonifacioJZ/todo/db"
	"github.com/BonifacioJZ/todo/handlers"
	"github.com/BonifacioJZ/todo/migrations"
)

func main() {

	if db.CheckConnect() == 0 {

		log.Fatal("Error en la base de datos")
	}
	migrations.Initmigrations()
	handlers.Handlers()

}
