package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable          = "users"
	todoListsTable      = "todo_lists"
	usersTodoListsTable = "users_lists"
	todoItemTable       = "todo_item"
	listsItemsTable     = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.ToString())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (c Config) ToString() string {
	return fmt.Sprintf("host=%s port=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, //c.Username,
		c.Password, c.DBName, c.SSLMode)
}
