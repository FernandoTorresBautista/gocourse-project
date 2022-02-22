package main

import (
	"log"

	"github.com/FernandoTorresBautista/gocourse-project/db"
	"github.com/FernandoTorresBautista/gocourse-project/handlers"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexión a DB")
	}
	handlers.Handlers()

}
