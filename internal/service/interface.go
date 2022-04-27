package service

import "context"

type Service interface {
	Create(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, url string) (string, error)

	CacheInit(ctx context.Context) error
}
