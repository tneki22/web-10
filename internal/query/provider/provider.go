package provider

import (
	"database/sql"
	"fmt"
	"log"

	"web-10/internal/query/model"

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

	// Создаем таблицу users, если ее нет
	_, err = conn.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL
        );
    `)
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{conn: conn}
}

func (p *Provider) GetUser(name string) (*model.User, error) {
	var user model.User
	err := p.conn.QueryRow(
		"SELECT id, name FROM users WHERE name = $1", name).Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *Provider) AddUser(name string) error {
	_, err := p.conn.Exec("INSERT INTO users (name) VALUES ($1)", name)
	return err
}
