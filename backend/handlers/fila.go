package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

type FilaHandler struct {
	db *sql.DB
}

func NewFilaHandler(database *sql.DB) *FilaHandler {
	return &FilaHandler{db: database}
}

func (h FilaHandler) CriarEntrada(c fiber.Ctx) error {
	return c.SendString("Registra a chegada de um motorista na fila ")
}

func (h FilaHandler) AtualizarStatus(c fiber.Ctx) error {
	return c.SendString("Atualiza o status de uma entrada na fila")
}

func (h FilaHandler) ListarEntradas(c fiber.Ctx) error {
	return c.SendString("Lista todos os motoristas na fila do dia, ordenados por chegada")
}

func (h FilaHandler) BuscarEntrada(c fiber.Ctx) error {
	return c.SendString("Retorna os detalhes de uma entrada específica")
}

func (h FilaHandler) Historico(c fiber.Ctx) error {
	return c.SendString("Lista entradas finalizadas e canceladas (dias anteriores)")
}
