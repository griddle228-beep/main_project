package app

import (
	"context"
	"fmt"
	"log/slog"
	"log"
	"os"
	"path/filepath"

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

	migrationsPath := "migrations"
	
	// Проверяем, существует ли папка
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		// Если не нашли, пробуем подняться выше
		migrationsPath = filepath.Join("..", "migrations")
		
		// Проверяем снова
		if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
			slog.Error("migrations directory not found", "tried", []string{"migrations", "../migrations"})
			return fmt.Errorf("migrations directory not found")
		}
	}
	slog.Info("found migrations", "path", migrationsPath)

	migrator := database.NewMigrator(dbPool)
    if err := migrator.RunMigrations(migrationsPath); err != nil {
        log.Fatal("Failed to run migrations:", err)
    }
	
	handler := controllers.NewHandlers(dbPool) 

	router := gin.Default()
	routes.SetupRoutes(router, handler)
	router.Run(fmt.Sprintf(":%d", cfg.PublicApiPort))
	return nil
}
