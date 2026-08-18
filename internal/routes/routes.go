package routes

import (
	"github.com/aditya-singh-finbox/todo-api/internal/auth"
	"github.com/aditya-singh-finbox/todo-api/internal/handler"
	"github.com/aditya-singh-finbox/todo-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authHandler *handler.AuthHandler, todoHandler *handler.TodoHandler, jwtService *auth.JWTService) {
	api := router.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.POST("/refresh", authHandler.Refresh)

		api.POST("/logout", authHandler.Logout)
		todos := api.Group("/todos")
		todos.Use(middleware.AuthMiddleware(jwtService))
		{
			todos.POST("", todoHandler.Create)
			todos.GET("", todoHandler.GetAll)
			todos.GET("/:id", todoHandler.GetByID)
			todos.PUT("/:id", todoHandler.Update)
			todos.DELETE("/:id", todoHandler.Delete)
		}
	}
}
