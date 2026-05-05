package main

import (
	"log"
	"fuel-terminal/db"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate(database)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MIGRADO!")
}
