package repositories_impl

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"uri-shortener/internal/services/repositories"
)

type linkRepositoryImpl struct {
	client *redis.Client
}

func NewLinkRepository(redisClient *redis.Client) repositories.LinkRepository {
	return &linkRepositoryImpl{
		client: redisClient,
	}
}

func (l *linkRepositoryImpl) Create(ctx context.Context, linkToShorten string, shortLink string, ttlMinutes int) (err error) {
	expiration := time.Duration(ttlMinutes) * time.Minute
	err = l.client.Set(ctx, shortLink, linkToShorten, expiration).Err()

	return err
}

func (l *linkRepositoryImpl) GetFullLink(ctx context.Context, shortLink string) (fullLink string, err error) {
	return l.client.Get(ctx, shortLink).Result()
}
