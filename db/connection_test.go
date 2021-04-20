package db_test

import (
	"testing"

	"github.com/BonifacioJZ/todo/db"
)

//Test Para Probar la Conexion ala base de datos
func TestConnectDB(t *testing.T) {

	conn := db.ConnectionDB()

	if conn == nil {
		t.Error("Error to connect db")
		t.Fail()
	} else {
		t.Log("Connection success")
	}

}

//Test para probar el ping ala base de datos

func TestPing(t *testing.T) {
	ping := db.CheckConnect()

	if ping == 0 {
		t.Error("Error to ping db")
		t.Fail()
	} else {
		t.Log("Ping success")
	}
}
