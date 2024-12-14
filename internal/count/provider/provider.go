package provider

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Provider struct {
	conn *sql.DB
}

func NewProvider(host string, port int, user, password, dbName string) *Provider {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	// Создаем таблицу counter, если ее нет
	_, err = conn.Exec(`
        CREATE TABLE IF NOT EXISTS counter (
            id SERIAL PRIMARY KEY,
            value INTEGER NOT NULL
        );
    `)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем счетчик, если записи нет
	var exists bool
	err = conn.QueryRow("SELECT EXISTS(SELECT 1 FROM counter)").Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		_, err = conn.Exec("INSERT INTO counter (value) VALUES (0)")
		if err != nil {
			log.Fatal(err)
		}
	}

	return &Provider{conn: conn}
}
