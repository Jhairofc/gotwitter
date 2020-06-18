package main

import (
	"log"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/handlers"
)

func main() {
	if db.CheckConnectionDB() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Handlers()
}
