package pgconn

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectToPg создает подключение к PostgreSQL
func ConnectToPg(ctx context.Context, cfg DbConfig, appName string) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(cfg.PgConnDsn())
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrWrongConfig, err.Error())
	}

	pgxConfig.MaxConns = int32(cfg.PgMaxConn)
	pgxConfig.ConnConfig.RuntimeParams["application_name"] = appName

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrEstablishConnect, err.Error())
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("%w: ping failed: %s", ErrEstablishConnect, err.Error())
	}

	return pool, nil
}

