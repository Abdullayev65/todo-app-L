package repository

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	UserIdBy(username, password_hash string) (int, error)
}
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId int, listId int, input todo.InputListUpdate) error
}
type TodoItem interface {
	Create(userId int, listId int, list todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, listId, itemId int) (todo.TodoItem, error)
	Delete(userId, listId, itemId int) error
	Update(userId int, listId int, itemId int, input todo.InputItemUpdate) error
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewItemPostgres(db),
	}
}
