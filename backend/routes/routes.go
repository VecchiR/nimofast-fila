package routes

import (
	"fuel-terminal/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	produto := handlers.NewProdutoHandler()

	api := app.Group("/api/v1")

	// Produtos
	api.Get("/produtos", produto.Listar)
}
