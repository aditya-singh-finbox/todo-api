package repository

import (
	"github.com/aditya-singh-finbox/todo-api/internal/database"
	"github.com/aditya-singh-finbox/todo-api/internal/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{db: database.GetDB()}
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}
func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) GetByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Update(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) Delete(id uint) error {
	return r.db.Delete(&model.Todo{}, id).Error
}
