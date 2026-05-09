package handlers

import (
	"database/sql"
	"log"
	"fuel-terminal/models"

	"github.com/gofiber/fiber/v3"
)

type ProdutoHandler struct {
	db *sql.DB
}

func NewProdutoHandler(database *sql.DB) *ProdutoHandler {
	return &ProdutoHandler{db: database}
}

func (h ProdutoHandler) Listar(c fiber.Ctx) error {
	query := "SELECT id, nome FROM produtos"
	rows, err := h.db.Query(query)
	if err != nil {
		log.Printf("[ERROR] Query failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Não foi possível buscar os produtos no momento."})
	}
	defer rows.Close()

	produtos := make([]models.Produto, 0)
	for rows.Next() {
		var p models.Produto
		err := rows.Scan(&p.ID, &p.Nome)
		if err != nil {
			log.Printf("[ERROR] Scan failed: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Erro ao processar a lista de produtos."})
		}
		produtos = append(produtos, p)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[ERROR] Row iteration failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar os produtos devido a uma falha de conexão."})
	}

	return c.JSON(produtos)
}
