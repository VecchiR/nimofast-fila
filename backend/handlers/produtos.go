package handlers

import "github.com/gofiber/fiber/v3"

type ProdutoHandler struct {
	// .... mas não sei o que colocar aqui :,)
}

func NewProdutoHandler() *ProdutoHandler {
	return &ProdutoHandler{}
}

func (h ProdutoHandler) Listar(c fiber.Ctx) error {
	return c.SendString("Lista os produtos disponíveis para carregamento")
}
