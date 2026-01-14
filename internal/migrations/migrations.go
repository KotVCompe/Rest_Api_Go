package migrations

import (
	"fmt"
	"go-api-example/internal/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	// Автоматическая миграция для создания таблиц
	err := db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		return fmt.Errorf("ошибка миграции: %w", err)
	}

	fmt.Println("Миграции успешно выполнены")
	return nil
}
