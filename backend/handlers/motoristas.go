package handlers

import (
	"database/sql"
	"fuel-terminal/models"
	"log"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type MotoristaHandler struct {
	db *sql.DB
}

func NewMotoristaHandler(database *sql.DB) *MotoristaHandler {
	return &MotoristaHandler{db: database}
}

var (
	cpfRegex   = regexp.MustCompile(`[^0-9]`)
	placaRegex = regexp.MustCompile(`[^A-Z0-9]`)
)

func (h MotoristaHandler) Buscar(c fiber.Ctx) error {

	cpf := c.Query("cpf")
	placa := strings.ToUpper(c.Query("placa"))

	if cpf == "" && placa == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Forneça um CPF ou a placa do veículo para a busca",
		})
	}

	if cpf != "" && placa != "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Forneça um único parâmetro para a busca: CPF ou placa do veículo",
		})
	}

	if cpf != "" {
		cpfLimpo := cpfRegex.ReplaceAllString(cpf, "")
		if len(cpfLimpo) != 11 {
			return c.Status(400).JSON(fiber.Map{
				"error": "CPF deve ser composto por 11 dígitos numéricos",
			})
		}
		var m models.Motorista

		query := `SELECT id, nome, cpf, cnh, placa FROM motoristas WHERE cpf = $1`
		err := h.db.QueryRow(query, cpfLimpo).Scan(&m.ID, &m.Nome, &m.CPF, &m.CNH, &m.Placa)
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Não foi encontrado um motorista através do CPF informado"})
		}
		if err != nil {
			log.Printf("[ERROR] Query failed: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar motorista"})
		}

		return c.JSON(m)

	}

	if placa != "" {
		placaLimpa := placaRegex.ReplaceAllString(placa, "")
		if len(placaLimpa) != 7 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Placa deve conter 7 caracteres alfanuméricos",
			})
		}
		var m models.Motorista

		query := `SELECT id, nome, cpf, cnh, placa FROM motoristas WHERE placa = $1`
		err := h.db.QueryRow(query, placaLimpa).Scan(&m.ID, &m.Nome, &m.CPF, &m.CNH, &m.Placa)
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Não foi encontrado um motorista através da placa informada"})
		}
		if err != nil {
			log.Printf("[ERROR] Query failed: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar motorista"})
		}

		return c.JSON(m)

	}

	log.Printf("[ERROR] Error during search: cpf=%s, placa=%s", cpf, placa)
	return c.Status(400).JSON(fiber.Map{
		"error": "Erro ao processar a busca",
	})

}
