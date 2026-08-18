package main

import (
	"log"

	"github.com/aditya-singh-finbox/todo-api/config"
	"github.com/aditya-singh-finbox/todo-api/internal/auth"
	"github.com/aditya-singh-finbox/todo-api/internal/database"
	"github.com/aditya-singh-finbox/todo-api/internal/handler"
	"github.com/aditya-singh-finbox/todo-api/internal/model"
	"github.com/aditya-singh-finbox/todo-api/internal/repository"
	"github.com/aditya-singh-finbox/todo-api/internal/routes"
	"github.com/aditya-singh-finbox/todo-api/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	database.Connect(cfg)

	db := database.GetDB()

	if err := db.AutoMigrate(
		&model.User{},
		&model.Todo{},
		model.RefreshToken{},
	); err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	todoRepository := repository.NewTodoRepository()
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)

	userService := service.NewUserService(userRepo)
	todoService := service.NewTodoService(todoRepository)
	jwtService := auth.NewJWTService(cfg.JWTSecret)

	authService := service.NewAuthService(
		userService,
		jwtService,
		refreshTokenRepo,
	)

	authHandler := handler.NewAuthHandler(
		userService,
		authService,
	)
	todoHandler := handler.NewTodoHandler(todoService)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Todo API is running",
		})
	})
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		AllowCredentials: true,
	}))
	routes.SetupRoutes(
		router,
		authHandler,
		todoHandler,
		jwtService,
	)
	log.Printf("Starting server on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
