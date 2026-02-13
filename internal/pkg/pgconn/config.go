package pgconn

import (
	"fmt"
)

type DbConfig struct {
	PgHost    string `env:"PG_HOST,notEmpty"`
	PgPort    uint16 `env:"PG_PORT,notEmpty"`
	PgUser    string `env:"PG_USER,notEmpty"`
	PgPass    string `env:"PG_PASS,notEmpty" logging:"mask"`
	PgDbName  string `env:"PG_DB,notEmpty"`
	PgMaxConn uint16 `env:"PG_MAX_CONN" envDefault:"20"`
}

func (cfg *DbConfig) PgConnDsn() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cfg.PgUser, cfg.PgPass, cfg.PgHost, cfg.PgPort, cfg.PgDbName,
	)
}
