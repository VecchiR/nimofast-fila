package main

import (
	"fuel-terminal/routes"

	"github.com/gofiber/fiber/v3"
)

// func main() {
// 	database, err := db.Connect()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = db.Migrate(database)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("MIGRADO!")
// }

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)
	app.Listen(":3000")
}
