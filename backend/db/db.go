package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func Connect() {
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
	defer db.Close()

	err = db.Ping()
	if err == nil {
		log.Println("CONECTADO!")
	} else {
		log.Fatal(err)
	}
}
