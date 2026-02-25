package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"semen_project/internal/routes"
	"semen_project/internal/pkg/pgconn"
	"semen_project/internal/config"
	"semen_project/internal/controllers"
	"github.com/gin-gonic/gin"
	"semen_project/internal/database"
)

// Run основная функция запуска приложения
func Run(cfg *config.Config) error {
	ctx := context.Background() 
	slog.Info("initializing application", "app_name", cfg.AppName )

	// Подключаемся к PostgreSQL
	slog.Info("connecting to PostgreSQL database...")

	dbPool, err := pgconn.ConnectToPg(ctx, cfg.PG, cfg.AppName) 
	if err != nil {
		slog.Error("failed to connect to PostgreSQL", "error", err)
		return fmt.Errorf("database connection failed: %w", err)
	}
	defer dbPool.Close()

	slog.Info("PostgreSQL connection established")

	// ИСПРАВЛЕННЫЙ КОД С ЛОГИРОВАНИЕМ
	// ============================================
	
	// Получаем текущую директорию
	currentDir, _ := os.Getwd()
	slog.Info("current directory", "path", currentDir)
	
	// Правильный путь к миграциям (в корне проекта)
	migrationsPath := "migrations"
	
	// Проверяем существование папки
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		slog.Error("migrations folder not found!", "path", migrationsPath)
		return fmt.Errorf("migrations folder not found at %s", migrationsPath)
	}
	
	// Проверяем содержимое папки
	files, _ := os.ReadDir(migrationsPath)
	slog.Info("migrations folder content:")
	for _, f := range files {
		if !f.IsDir() {
			slog.Info("  - " + f.Name())
		}
	}
	
	// СОЗДАЕМ МИГРАТОР
	migrator := database.NewMigrator(dbPool)
	slog.Info("starting migrations...")
	
	// ЗАПУСКАЕМ МИГРАЦИИ - ЭТО КЛЮЧЕВОЙ МОМЕНТ!
	err = migrator.RunMigrations(migrationsPath)
	if err != nil {
		slog.Error("failed to run migrations", "error", err)
		return fmt.Errorf("migration failed: %w", err)
	}
	
	slog.Info("migrations completed successfully")
	
	// Проверяем, создалась ли таблица
	var tableExists bool
	checkQuery := `SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'users')`
	err = dbPool.QueryRow(ctx, checkQuery).Scan(&tableExists)
	if err != nil {
		slog.Error("failed to check table existence", "error", err)
	} else if tableExists {
		slog.Info("✅ Table 'users' exists")
	} else {
		slog.Error("❌ Table 'users' does NOT exist")
	}
	
	handler := controllers.NewHandlers(dbPool) 

	router := gin.Default()
	routes.SetupRoutes(router, handler)
	router.Run(fmt.Sprintf(":%d", cfg.PublicApiPort))
	return nil
}
