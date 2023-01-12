package service

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/repository"
)

type ItemService struct {
	repo repository.TodoItem
}

func NewItemService(repo repository.TodoItem) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) Create(userId int, listId int, list todo.TodoItem) (int, error) {
	return s.repo.Create(userId, listId, list)
}

func (s *ItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *ItemService) GetById(userId, listId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, listId, itemId)
}

func (s *ItemService) Delete(userId, listId, itemId int) error {
	return s.repo.Delete(userId, listId, itemId)
}
func (s *ItemService) Update(userId int, listId int, itemId int, input todo.InputListUpdate) error {
	return s.repo.Update(userId, listId, itemId, input)
}
