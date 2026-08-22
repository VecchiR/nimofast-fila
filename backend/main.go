package main

import (
	"fuel-terminal/db"
	"fuel-terminal/routes"
	"log"
	"os"

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
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}
	cfg.AllowOrigins = []string{frontendURL}
	app.Use(cors.New(cfg))

	routes.SetupRoutes(app, database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Println("Iniciando server em ", addr)
	err = app.Listen(addr)
	if err != nil {
		log.Fatal(err)
	}

}
