package routes

import (
	"fuel-terminal/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	produto := handlers.NewProdutoHandler()
	motorista := handlers.NewMotoristaHandler()
	fila := handlers.NewFilaHandler()

	api := app.Group("/api/v1")

	// Produtos
	api.Get("/produtos", produto.Listar)

	// Motoristas
	api.Get("/motoristas/busca", motorista.Buscar)

	// Fila
	api.Get("/fila/historico", fila.Historico)
	api.Get("/fila/:id", fila.BuscarEntrada)
	api.Get("/fila", fila.ListarEntradas)
	api.Patch("/fila/:id/status", fila.AtualizarStatus)
	api.Post("/fila", fila.CriarEntrada)
}
