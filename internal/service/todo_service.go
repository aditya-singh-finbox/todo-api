package service

import (
	"errors"
	"strings"

	"github.com/aditya-singh-finbox/todo-api/internal/model"
	"github.com/aditya-singh-finbox/todo-api/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) Create(todo *model.Todo) error {
	todo.Title = strings.TrimSpace(todo.Title)
	if todo.Title == "" {
		return errors.New("title cannot be empty")
	}
	if len(todo.Title) > 100 {
		return errors.New("title cannot exceed 100 characters")
	}
	return s.repo.Create(todo)
}

func (s *TodoService) GetAll() ([]model.Todo, error) {
	return s.repo.GetAll()
}
func (s *TodoService) GetByID(id uint) (*model.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *TodoService) Update(todo *model.Todo) error {
	todo.Title = strings.TrimSpace(todo.Title)
	if todo.Title == "" {
		return errors.New("title cannot be empty")
	}
	if len(todo.Title) > 100 {
		return errors.New("title cannot exceed 100 characters")
	}
	return s.repo.Update(todo)
}

func (s *TodoService) Delete(id uint) error {
	return s.repo.Delete(id)
}
