package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

type MotoristaHandler struct {
	db *sql.DB
}

func NewMotoristaHandler(database *sql.DB) *MotoristaHandler {
	return &MotoristaHandler{db: database}
}

func (h MotoristaHandler) Buscar(c fiber.Ctx) error {
	return c.SendString("Busca motorista por CPF ou placa para pré-preenchimento do formulário ")
}
