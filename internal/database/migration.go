package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Migrator struct {
	db *pgxpool.Pool
}

func NewMigrator(db *pgxpool.Pool) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) RunMigrations(migrationsDir string) error {
	ctx := context.Background()
	// Читаем все .up.sql файлы из папки
	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.up.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	for _, file := range files {
		log.Printf("Applying migration: %s", filepath.Base(file))

		// Читаем содержимое файла
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}
		// Выполняем SQL
		_, err = m.db.Exec(ctx, string(content))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}

		log.Printf("✅ Migration %s applied", filepath.Base(file))
	}

	return nil
}
