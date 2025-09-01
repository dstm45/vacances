// Package database gère la connexion à la base de données et les requêtes.
package database

import (
	"context"
	"time"

	"github.com/dstm45/vacances/pkg/api"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func Connection(ctx context.Context, config api.Config) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(config.Dsn)
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = 1000
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = time.Minute * 15
	poolConfig.HealthCheckPeriod = time.Minute
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
