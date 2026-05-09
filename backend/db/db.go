package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	// Capture connection properties.
	cfg := pq.Config{
		Host:           "localhost",
		Port:           5433,
		User:           "postgres",
		Password:       "postgres",
		Database:       "fuel_terminal",
		ConnectTimeout: time.Second * 10,
		SSLMode:        "disable",
	}

	connector, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db := sql.OpenDB(connector)

	err = db.Ping()
	if err == nil {
		log.Println("CONECTADO!")
	} else {
		log.Fatal(err)
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS motoristas (
			id         SERIAL PRIMARY KEY,
			nome       VARCHAR(100) NOT NULL,
			cpf        CHAR(11) UNIQUE NOT NULL,
			cnh        CHAR(9) UNIQUE NOT NULL,
			placa      CHAR(7) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
			CONSTRAINT check_placa_uppercase CHECK (placa = UPPER(placa))
		);

		CREATE TABLE IF NOT EXISTS produtos (
			id    SERIAL PRIMARY KEY,
			nome  VARCHAR(100) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS entradas_fila (
			id                   SERIAL PRIMARY KEY,
			motorista_id         INT NOT NULL REFERENCES motoristas(id),
			produto_id           INT NOT NULL REFERENCES produtos(id),
			status               VARCHAR(20) NOT NULL DEFAULT 'AGUARDANDO',
			horario_chegada      TIMESTAMP NOT NULL DEFAULT NOW(),
			inicio_carregamento  TIMESTAMP,
			fim_carregamento     TIMESTAMP
		);

		-- seed produtos (apenas se a tabela estiver vazia)
		INSERT INTO produtos (nome)
		SELECT unnest(ARRAY[
			'Diesel S10 A',
			'Diesel S10 B',
			'Diesel S10 B Aditivado',
			'Diesel S500 A',
			'Diesel S500 B',
			'Diesel S500 B Aditivado',
			'Gasolina A',
			'Gasolina C',
			'Gasolina C Aditivada'
		])
		WHERE NOT EXISTS (SELECT 1 FROM produtos);
	`)

	return err
}
