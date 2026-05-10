package models

import (
	"time"
)

type Status string

const (
	StatusAguardando Status = "AGUARDANDO"
	StatusCarregando Status = "CARREGANDO"
	StatusFinalizado Status = "FINALIZADO"
	StatusCancelado  Status = "CANCELADO"
)

func (atual Status) PodeTrocar(novo Status) bool {
	switch atual {
	case StatusAguardando:
		return novo == StatusCarregando || novo == StatusCancelado
	case StatusCarregando:
		return novo == StatusFinalizado || novo == StatusCancelado
	case StatusFinalizado, StatusCancelado:
		return false
	default:
		return false
	}
}

func (atual Status) IsStatusValido() bool {
	return atual == StatusAguardando || atual == StatusCarregando || atual == StatusFinalizado || atual == StatusCancelado
}

type Motorista struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	CPF   string `json:"cpf"`
	CNH   string `json:"cnh"`
	Placa string `json:"placa"`
}

type Produto struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

type EntradaFila struct {
	ID                 int        `json:"id"`
	MotoristaID        int        `json:"motorista_id"`
	Motorista          *Motorista `json:"motorista,omitempty"`
	ProdutoID          int        `json:"produto_id"`
	Produto            *Produto   `json:"produto,omitempty"`
	Status             Status     `json:"status"`
	HorarioChegada     time.Time  `json:"horario_chegada"`
	InicioCarregamento *time.Time `json:"inicio_carregamento,omitempty"`
	FimCarregamento    *time.Time `json:"fim_carregamento,omitempty"`
}
