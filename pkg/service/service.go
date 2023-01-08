package service

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateTokenIfExists(username, password string) (string, error)
}
type TodoList interface{}
type TodoItem interface{}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repo),
	}
}
