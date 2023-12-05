package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(pc *PostgresConfig) (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		pc.Username, pc.Password, pc.Host, pc.Port, pc.DBName, pc.SSLMode,
	)

	db, err := pgx.Connect(ctx, url)
	if err != nil {
		return db, err
	}
	
	return db, nil
}
