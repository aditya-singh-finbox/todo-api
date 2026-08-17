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
func (r *TodoRepository) GetAllByuserID(userID uint) ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Where("user_id=?", userID).Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) GetByID(userID uint, todoID uint) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.Where("id =? AND user_id =?", todoID, userID).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Update(todo *model.Todo) error {
	return r.db.Model(&model.Todo{}).Where("id = ? AND user_id = ?", todo.ID, todo.UserID).Updates(todo).Error
}

func (r *TodoRepository) Delete(userId uint, todoID uint) error {
	return r.db.Where("id = ? AND user_id = ?", todoID, userId).Delete(&model.Todo{}).Error
}
