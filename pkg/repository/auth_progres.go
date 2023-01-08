package repository

import (
	"fmt"
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func newAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s(name, username, password_hash) VALUES($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *AuthPostgres) UserIdBy(username, passwordHash string) (int, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	row := r.db.QueryRow(query, username, passwordHash)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
