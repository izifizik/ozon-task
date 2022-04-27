package service

import (
	"context"
	"fmt"
	"math/rand"
	"ozon-task/internal/repository/cache"
	"ozon-task/internal/repository/database"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type service struct {
	alphabet string
	repo     database.Repository
	cache    cache.Cache
}

func NewService(r database.Repository, c cache.Cache) Service {
	a := "qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM1234567890_"
	return &service{repo: r, alphabet: a, cache: c}
}

func (s *service) Create(ctx context.Context, url string) (string, error) {
	shortUrl := "http://localhost:8080/" + s.generateShortUrl()
	return shortUrl, s.repo.Create(ctx, url, shortUrl)
}

func (s *service) Get(ctx context.Context, shortUrl string) (string, error) {
	url, ok := s.cache.Get(shortUrl)
	if !ok {
		return "", fmt.Errorf("error with get url by short url (empty)")
	}

	return s.repo.Get(ctx, url)
}

func (s *service) CacheInit(ctx context.Context) error {
	return s.getSQLRows(ctx)
}

func (s *service) getSQLRows(ctx context.Context) error {
	rows, err := s.repo.GetAll(ctx)
	if err != nil {
		return err
	}

	var url, shortUrl string

	for rows.Next() {
		err = rows.Scan(&url, shortUrl)
		if err != nil {
			return err
		}
		s.cache.Set(url, shortUrl)
	}
	return nil
}

func (s *service) generateShortUrl() string {
	url := ""
	max := len(s.alphabet)
	min := 0

	for i := 0; i < 10; i++ {
		r := rand.Intn(max-min) + min
		url += string(s.alphabet[r])
	}

	return url
}
