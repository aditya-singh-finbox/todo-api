package service

import (
	"errors"
	"strings"

	"github.com/aditya-singh-finbox/todo-api/internal/model"
	"github.com/aditya-singh-finbox/todo-api/internal/repository"
	"gorm.io/gorm"
)

type TodoService struct {
	todoRepo *repository.TodoRepository
}

func NewTodoService(todoRepo *repository.TodoRepository) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
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
	return s.todoRepo.Create(todo)
}

func (s *TodoService) GetAll(userID uint) ([]model.Todo, error) {
	return s.todoRepo.GetAllByuserID(userID)
}
func (s *TodoService) GetByID(userID uint, todoID uint) (*model.Todo, error) {
	return s.todoRepo.GetByID(userID, todoID)
}

func (s *TodoService) Update(userID uint, todoID uint, updatedTodo *model.Todo) error {
	existingTodo, err := s.todoRepo.GetByID(userID, todoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("todo not found")
		}
		return err
	}

	existingTodo.Title = updatedTodo.Title
	existingTodo.Description = updatedTodo.Description
	existingTodo.Completed = updatedTodo.Completed

	return s.todoRepo.Update(existingTodo)

}

func (s *TodoService) Delete(userID uint, todoID uint) error {

	err := s.todoRepo.Delete(userID, todoID)

	if err != nil {
		return err
	}
	return nil
}
