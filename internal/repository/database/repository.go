package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

type repo struct {
	db *pgx.Conn
}

func NewRepo(db *pgx.Conn) Repository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, url, shortUrl string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, "insert into test (url, short_url) values ($1, $2)", url, shortUrl)
	if err != nil {
		log.Println("Repo. Create." + err.Error())
		tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}

func (r *repo) Get(ctx context.Context, shortUrl string) (string, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return "", err
	}

	var url string
	err = r.db.QueryRow(ctx, "select (url) from test where short_url=$1", shortUrl).Scan(&url)
	if err != nil {
		log.Println("Repo. Get." + err.Error())
		tx.Rollback(ctx)
		return "", err
	}

	return url, tx.Commit(ctx)
}

func (r *repo) GetAll(ctx context.Context) (pgx.Rows, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, "select (url, short_url) from test")
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	return rows, tx.Commit(ctx)
}
