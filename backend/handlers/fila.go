package handlers

import (
	"database/sql"
	"log"
	"fuel-terminal/models"
	"strings"

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
		log.Printf("[ERROR] Failed to parse request body: %v", err)
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
			placa = EXCLUDED.placa
		RETURNING id
	`, req.Nome, req.CPF, req.CNH, req.Placa).Scan(&motoristaID)
	if err != nil {
		log.Printf("[ERROR] Failed to insert or update motorista: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao cadastrar ou atualizar motorista"})
	}

	// Verificar se o motorista já possui uma entrada ativa na fila (não pode ter mais de uma entrada ativa simultaneamente)
	var count int
	if err := h.db.QueryRow(`
		SELECT COUNT(*) FROM entradas_fila
		WHERE motorista_id = $1 AND status IN ('AGUARDANDO', 'CARREGANDO')
	`, motoristaID).Scan(&count); err != nil {
		log.Printf("[ERROR] Scan failed: %v", err)
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
		log.Printf("[ERROR] Failed to create fila entry: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao criar entrada na fila"})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":      entradaID,
		"message": "Motorista adicionado à fila com sucesso",
	})
}

func (h FilaHandler) AtualizarStatus(c fiber.Ctx) error {

	id := fiber.Params[int](c, "id")
	if id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "O ID da entrada é obrigatório e deve ser um número inteiro"})
	}

	var reqBody struct {
		StatusNovo models.Status `json:"status_novo"`
	}
	if err := c.Bind().Body(&reqBody); err != nil {
		log.Printf("[ERROR] Failed to parse request body: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Erro ao processar os dados fornecidos pela requisição"})
	}

	reqBody.StatusNovo = models.Status(strings.ToUpper(string(reqBody.StatusNovo)))

	if !reqBody.StatusNovo.IsStatusValido() {
		log.Printf("[ERROR] Invalid status: %v", reqBody.StatusNovo)
		return c.Status(400).JSON(fiber.Map{"error": "O Status fornecido é inválido"})
	}

	var statusAtual models.Status
	err := h.db.QueryRow(`SELECT status FROM entradas_fila WHERE id = $1`, id).Scan(&statusAtual)
	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{"error": "Entrada não encontrada"})
	}
	if err != nil {
		log.Printf("[ERROR] Failed to query current status: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao consultar entrada"})
	}

	if !statusAtual.PodeTrocar(reqBody.StatusNovo) {
		return c.Status(400).JSON(fiber.Map{"error": "Transição de status inválida"})
	}

	var updateQuery string
	switch reqBody.StatusNovo {
	case models.StatusCarregando:
		updateQuery = `UPDATE entradas_fila SET status = $1, inicio_carregamento = NOW() WHERE id = $2`
	case models.StatusFinalizado:
		updateQuery = `UPDATE entradas_fila SET status = $1, fim_carregamento = NOW() WHERE id = $2`
	default:
		// cancelado
		updateQuery = `UPDATE entradas_fila SET status = $1 WHERE id = $2`
	}

	_, err = h.db.Exec(updateQuery, reqBody.StatusNovo, id)
	if err != nil {
		log.Printf("[ERROR] Failed to update status: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao atualizar status da entrada"})
	}

	return c.JSON(fiber.Map{"message": "Status atualizado com sucesso"})
}

func (h FilaHandler) ListarEntradas(c fiber.Ctx) error {
	query := `SELECT id, motorista_id, produto_id, status, horario_chegada, inicio_carregamento, fim_carregamento FROM entradas_fila WHERE horario_chegada >= CURRENT_DATE ORDER BY horario_chegada ASC`

	rows, err := h.db.Query(query)
	if err != nil {
		log.Printf("[ERROR] Query failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao consultar entradas da fila"})
	}
	defer rows.Close()

	entradas := make([]models.EntradaFila, 0)
	for rows.Next() {
		var e models.EntradaFila
		if err := rows.Scan(&e.ID, &e.MotoristaID, &e.ProdutoID, &e.Status, &e.HorarioChegada, &e.InicioCarregamento, &e.FimCarregamento); err != nil {
			log.Printf("[ERROR] Scan failed: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Erro ao processar entradas da fila"})
		}
		entradas = append(entradas, e)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[ERROR] Row iteration failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar entradas da fila devido a uma falha de conexão"})
	}

	return c.JSON(entradas)

}

func (h FilaHandler) BuscarEntrada(c fiber.Ctx) error {

	entradaID := fiber.Params[int](c, "id")
	if entradaID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "O ID da entrada é obrigatório e deve ser um número inteiro"})
	}

	var e models.EntradaFila

	query := `SELECT id, motorista_id, produto_id, status, horario_chegada, inicio_carregamento, fim_carregamento FROM entradas_fila WHERE id = $1`
	if err := h.db.QueryRow(query, entradaID).Scan(&e.ID, &e.MotoristaID, &e.ProdutoID, &e.Status, &e.HorarioChegada, &e.InicioCarregamento, &e.FimCarregamento); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Entrada não encontrada"})
		}
		log.Printf("[ERROR] Scan failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar entrada"})
	}

	return c.JSON(e)
}

func (h FilaHandler) ListarHistorico(c fiber.Ctx) error {
	query := `SELECT id, motorista_id, produto_id, status, horario_chegada, inicio_carregamento, fim_carregamento FROM entradas_fila WHERE horario_chegada < CURRENT_DATE AND status in ('FINALIZADO', 'CANCELADO') ORDER BY horario_chegada DESC`

	rows, err := h.db.Query(query)
	if err != nil {
		log.Printf("[ERROR] Query failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao consultar histórico de entradas"})
	}
	defer rows.Close()

	entradas := make([]models.EntradaFila, 0)
	for rows.Next() {
		var e models.EntradaFila
		if err := rows.Scan(&e.ID, &e.MotoristaID, &e.ProdutoID, &e.Status, &e.HorarioChegada, &e.InicioCarregamento, &e.FimCarregamento); err != nil {
			log.Printf("[ERROR] Scan failed: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Erro ao processar histórico de entradas"})
		}
		entradas = append(entradas, e)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[ERROR] Row iteration failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar histórico de entradas devido a uma falha de conexão"})
	}

	return c.JSON(entradas)
}
