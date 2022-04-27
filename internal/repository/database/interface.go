package database

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	Create(ctx context.Context, url, shortUrl string) error
	Get(ctx context.Context, url string) (string, error)
	GetAll(ctx context.Context) (pgx.Rows, error)
}
