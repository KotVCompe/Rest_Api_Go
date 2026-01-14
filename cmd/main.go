package main

import (
	"fmt"
	"go-api-example/internal/config"
	"go-api-example/internal/database"
	"go-api-example/internal/handlers"
	"go-api-example/internal/migrations"
	"go-api-example/internal/repositories"
	"go-api-example/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Подключение к базе данных
	if err := database.ConnectDB(cfg); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Выполнение миграций
	db := database.GetDB()
	if err := migrations.RunMigrations(db); err != nil {
		log.Fatalf("Ошибка миграций: %v", err)
	}

	// Инициализация зависимостей
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Настройка роутера
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Группа маршрутов API
	api := router.Group("/api")
	{
		// Пользователи
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":   "ok",
				"database": "connected",
			})
		})
	}

	// Запуск сервера
	fmt.Printf("Сервер запущен на порту %s\n", cfg.ServerPort)
	if err := router.Run(cfg.ServerPort); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
