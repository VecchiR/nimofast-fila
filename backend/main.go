package main

import (
	"log"
	"fuel-terminal/db"
	"fuel-terminal/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate(database)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("MIGRADO!")
	}

	app := fiber.New()

	routes.SetupRoutes(app, database)
	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
