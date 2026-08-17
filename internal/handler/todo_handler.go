package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aditya-singh-finbox/todo-api/internal/model"
	"github.com/aditya-singh-finbox/todo-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoHandler struct {
	todoService *service.TodoService
}

func NewTodoHandler(todoService *service.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

// Create Todo
func (h *TodoHandler) Create(c *gin.Context) {

	userID := c.GetUint("userID")

	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	// IMPORTANT:
	// Never trust user_id coming from the request.
	// Get it from the JWT instead.
	todo.UserID = userID

	if err := h.todoService.Create(&todo); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, todo)
}

// Get all user's todos
func (h *TodoHandler) GetAll(c *gin.Context) {

	userID := c.GetUint("userID")

	todos, err := h.todoService.GetAll(userID)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch todos",
		})

		return
	}

	c.JSON(http.StatusOK, todos)
}

// Get user's todo by ID
func (h *TodoHandler) GetByID(c *gin.Context) {

	userID := c.GetUint("userID")

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid todo id",
		})

		return
	}

	todo, err := h.todoService.GetByID(
		userID,
		uint(id),
	)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			c.JSON(http.StatusNotFound, gin.H{
				"error": "todo not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch todo",
		})

		return
	}

	c.JSON(http.StatusOK, todo)
}

// Update Todo
func (h *TodoHandler) Update(c *gin.Context) {

	userID := c.GetUint("userID")

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid todo id",
		})

		return
	}

	var request model.Todo

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	err = h.todoService.Update(
		userID,
		uint(id),
		&request,
	)

	if err != nil {

		if err.Error() == "todo not found" {

			c.JSON(http.StatusNotFound, gin.H{
				"error": "todo not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update todo",
		})

		return
	}

	todo, err := h.todoService.GetByID(
		userID,
		uint(id),
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch updated todo",
		})

		return
	}

	c.JSON(http.StatusOK, todo)
}

// Delete Todo
func (h *TodoHandler) Delete(c *gin.Context) {

	userID := c.GetUint("userID")

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid todo id",
		})

		return
	}

	err = h.todoService.Delete(
		userID,
		uint(id),
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete todo",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo deleted successfully",
	})
}
