package main

import (
	"log"
	"fuel-terminal/db"
	"fuel-terminal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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

	cfg := cors.ConfigDefault
	cfg.AllowOrigins = []string{"http://localhost:3000"}
	app.Use(cors.New(cfg))

	routes.SetupRoutes(app, database)
	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
