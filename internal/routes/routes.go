package routes

import (
	"github.com/aditya-singh-finbox/todo-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authHandler *handler.AuthHandler, todoHandler *handler.TodoHandler) {
	api := router.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		todos := api.Group("/todos")
		{
			todos.POST("", todoHandler.Create)
			todos.GET("", todoHandler.GetAll)
			todos.GET("/:id", todoHandler.GetByID)
			todos.PUT("/:id", todoHandler.Update)
			todos.DELETE("/:id", todoHandler.Delete)
		}
	}
}
