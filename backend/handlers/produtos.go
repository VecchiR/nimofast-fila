package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

type ProdutoHandler struct {
	db *sql.DB
}

func NewProdutoHandler(database *sql.DB) *ProdutoHandler {
	return &ProdutoHandler{db: database}
}

func (h ProdutoHandler) Listar(c fiber.Ctx) error {
	return c.SendString("Lista os produtos disponíveis para carregamento")
}
