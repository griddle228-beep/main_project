package controllers

import (
	"semen_project/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)
type Handlers struct {
	DbPool *repository.Store
	Secret string
}

func NewHandlers(dbPool *pgxpool.Pool, secret string) *Handlers {
	return &Handlers{DbPool: repository.NewStore(dbPool), Secret: secret}
}