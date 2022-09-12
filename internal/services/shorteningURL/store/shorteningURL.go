package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

const CacheDuration = 6 * time.Hour

type Store struct {
	db *redis.Client
}

func New(db *redis.Client) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) SaveUrlMapping(ctx context.Context, shortUrl string, originalUrl string) (string, error) {
	res, err := s.db.Set(ctx, shortUrl, originalUrl, CacheDuration).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *Store) RetrieveInitialUrl(ctx context.Context, shortUrl string) (string, error) {
	res, err := s.db.Get(ctx, shortUrl).Result()
	if err != nil {
		return res, err
	}
	return res, nil
}
