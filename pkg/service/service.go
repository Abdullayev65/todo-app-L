package service

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateTokenIfExists(username, password string) (string, error)
	ParseToken(tokenStr string) (int, error)
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
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewItemService(repo.TodoItem),
	}
}
