package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var psql = ConnectionDB()

//Funcion para Conectarse a la base de datos
func ConnectionDB() *gorm.DB {
	var (
		host     = "localhost"
		user     = "reb"
		password = "ghost699"
		name     = "reb"
	)
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, name, password)

	Conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error in connect the DB %v", err)
		return nil
	}
	if Conn.Error != nil {
		log.Fatalln("Any Error in connect the DB " + err.Error())
		return nil
	}
	log.Println("DB connected")
	return Conn
}

//Funcion para Saber el estado de la base de datos
func CheckConnect() int {
	db, _ := psql.DB()

	if db.Ping() == nil {
		return 1
	} else {
		return 0
	}
}
