package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

type ProdutoHandler struct {
	// .... mas não sei o que colocar aqui :,)
}

func NewProdutoHandler(database *sql.DB) *ProdutoHandler {
	return &ProdutoHandler{}
}

func (h ProdutoHandler) Listar(c fiber.Ctx) error {
	return c.SendString("Lista os produtos disponíveis para carregamento")
}
