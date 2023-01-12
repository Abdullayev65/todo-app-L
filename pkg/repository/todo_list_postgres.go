package repository

import (
	"errors"
	"fmt"
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(query, list.Title, list.Description)
	var listId int
	if err = row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}
	usersListQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES ($1, $2)", usersTodoListsTable)
	_, err = tx.Exec(usersListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return listId, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	query := fmt.Sprintf("SELECT l.* FROM %s l INNER JOIN %s ul ON l.id=list_id WHERE user_id = $1", todoListsTable, usersTodoListsTable)
	var lists []todo.TodoList
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	query := fmt.Sprintf("SELECT l.* FROM %s l INNER JOIN %s ul ON l.id=list_id WHERE user_id = $1 AND list_id = $2", todoListsTable, usersTodoListsTable)
	var list todo.TodoList
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}
func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE ul.user_id=$1 AND tl.id = $2  AND tl.id=ul.list_id",
		todoListsTable, usersTodoListsTable)
	_, err := r.db.Exec(query, userId, listId)
	return err
}
func (r *TodoListPostgres) Update(userId int, listId int, input todo.InputListUpdate) error {
	argNum, args, setValues := 1, *new([]any), *new([]string)
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argNum))
		args = append(args, *input.Title)
		argNum++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argNum))
		args = append(args, *input.Description)
		argNum++
	}
	if argNum == 1 {
		return errors.New("all fields are nil")
	}
	setStr := strings.Join(setValues, " ,")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s lu WHERE lu.user_id= $%d AND tl.id = $%d AND lu.list_id = tl.id",
		todoListsTable, setStr, usersTodoListsTable, argNum, argNum+1)
	args = append(args, userId, listId)
	_, err := r.db.Exec(query, args...)
	return err
}
