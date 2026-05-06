package handlers

import "github.com/gofiber/fiber/v3"

type MotoristaHandler struct {
	// .... mas não sei o que colocar aqui :,)
}

func NewMotoristaHandler() *FilaHandler {
	return &FilaHandler{}
}

func (h FilaHandler) Buscar(c fiber.Ctx) error {
	return c.SendString("Busca motorista por CPF ou placa para pré-preenchimento do formulário ")
}
