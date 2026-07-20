package controllers

import (
	"semen_project/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)
type Handlers struct {
	DbPool *repository.Store
}

func NewHandlers(dbPool *pgxpool.Pool) *Handlers {
	return &Handlers{DbPool: repository.NewStore(dbPool)}
}