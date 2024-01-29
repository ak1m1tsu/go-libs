// Package postgresql provides a PostgreSQL connection.
package postgresql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connection represents a connection to a PostgreSQL database.
type Connection struct {
	cfg *Config

	*pgxpool.Pool
}

// Connect creates a new PostgreSQL connection.
func Connect(url string, cfg *Config) (*Connection, error) {
	if cfg == nil {
		cfg = NewConfig()
	}

	conn := &Connection{cfg: cfg}

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = conn.cfg.poolSize
	poolConfig.ConnConfig.ConnectTimeout = conn.cfg.connectionTimeout

	for attempt := uint(1); attempt < conn.cfg.connectionAttempts; attempt++ {
		conn.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			return conn, nil
		}

		if attempt < cfg.connectionAttempts {
			time.Sleep(cfg.retryDelay)
		}
	}

	return nil, err
}
