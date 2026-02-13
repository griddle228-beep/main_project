package app

import (
	"context"
	"fmt"
	"log/slog"

	"semen_project/internal/routes"
	"semen_project/internal/pkg/pgconn"
	"semen_project/internal/config"
	"semen_project/internal/controllers"
	"github.com/gin-gonic/gin"
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
	
	handler := controllers.NewHandlers(dbPool) 

	router := gin.Default()
	routes.SetupRoutes(router, handler)
	router.Run(fmt.Sprintf(":%d", cfg.PublicApiPort))
	return nil
}
