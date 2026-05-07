package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

type MotoristaHandler struct {
	// .... mas não sei o que colocar aqui :,)
}

func NewMotoristaHandler(database *sql.DB) *MotoristaHandler {
	return &MotoristaHandler{}
}

func (h MotoristaHandler) Buscar(c fiber.Ctx) error {
	return c.SendString("Busca motorista por CPF ou placa para pré-preenchimento do formulário ")
}
