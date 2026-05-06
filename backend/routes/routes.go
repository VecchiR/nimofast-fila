package routes

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.Get("/teste", func(c fiber.Ctx) error {
		return c.SendString("teste")
	})

	api.Get("/:teste", func(c fiber.Ctx) error {
		return c.SendString("teste parametro: " + c.Params("teste"))
	})
}
