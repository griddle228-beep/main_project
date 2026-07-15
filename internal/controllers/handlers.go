package controllers

import (
	"semen_project/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)
type Handlers struct {
	DbPool *database.Store
}

func NewHandlers(dbPool *pgxpool.Pool) *Handlers {
	return &Handlers{DbPool: database.NewStore(dbPool)}
}