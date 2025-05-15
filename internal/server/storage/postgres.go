package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"time"
)

type PostgresStorage struct {
	Conn *pgxpool.Pool
}

func NewPgStorage(ctx context.Context, dataBaseDSN string) (*PostgresStorage, error) {
	pool, err := pgxpool.New(ctx, dataBaseDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool, %w", err)
	}
	return &PostgresStorage{Conn: pool}, nil
}

func (p PostgresStorage) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := p.Conn.Ping(ctx); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (p PostgresStorage) IsUserExists() error {

	return nil
}
