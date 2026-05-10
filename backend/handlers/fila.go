package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v3"
)

type FilaHandler struct {
	db *sql.DB
}

func NewFilaHandler(database *sql.DB) *FilaHandler {
	return &FilaHandler{db: database}
}

type criarEntradaRequest struct {
	Nome      string `json:"nome"`
	CPF       string `json:"cpf"`
	CNH       string `json:"cnh"`
	Placa     string `json:"placa"`
	ProdutoID int    `json:"produto_id"`
}

func (h FilaHandler) CriarEntrada(c fiber.Ctx) error {
	var req criarEntradaRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Erro ao processar os dados fornecidos"})
	}

	// Verificar se todos os campos foram preenchidos
	if req.Nome == "" || req.CPF == "" || req.CNH == "" || req.Placa == "" || req.ProdutoID == 0 {
		log.Printf("[ERROR] Campos obrigatórios ausentes: %+v", req)
		return c.Status(400).JSON(fiber.Map{"error": "É necessário fornecer todos os campos: nome, cpf, cnh, placa e produto_id"})
	}

	// Criar ou atualizar os dados do motorista
	var motoristaID int
	err := h.db.QueryRow(`
		INSERT INTO motoristas (nome, cpf, cnh, placa)
		VALUES ($1, $2, $3, UPPER($4))
		ON CONFLICT (cpf) DO UPDATE SET
			nome = EXCLUDED.nome,
			cnh = EXCLUDED.cnh,
			placa = EXCLUDED.placa
		RETURNING id
	`, req.Nome, req.CPF, req.CNH, req.Placa).Scan(&motoristaID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao cadastrar ou atualizar motorista"})
	}

	// Verificar se o motorista já possui uma entrada ativa na fila (não pode ter mais de uma entrada ativa simultaneamente)
	var count int
	if err := h.db.QueryRow(`
		SELECT COUNT(*) FROM entradas_fila
		WHERE motorista_id = $1 AND status IN ('AGUARDANDO', 'CARREGANDO')
	`, motoristaID).Scan(&count); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao verificar entradas ativas do motorista"})
	}
	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Motorista já possui uma entrada ativa na fila"})
	}

	// Criar a entrada na fila
	var entradaID int
	if err := h.db.QueryRow(`
		INSERT INTO entradas_fila (motorista_id, produto_id, status, horario_chegada)
		VALUES ($1, $2, 'AGUARDANDO', NOW())
		RETURNING id
	`, motoristaID, req.ProdutoID).Scan(&entradaID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao criar entrada na fila"})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":      entradaID,
		"message": "Motorista adicionado à fila com sucesso",
	})
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
