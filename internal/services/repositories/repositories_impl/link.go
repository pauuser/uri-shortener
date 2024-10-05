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

func (l *linkRepositoryImpl) Create(ctx context.Context, linkToShorten string, shortLinkTail string, ttlMinutes int) (err error) {
	expiration := time.Duration(ttlMinutes) * time.Minute
	err = l.client.Set(ctx, shortLinkTail, linkToShorten, expiration).Err()

	return err
}

func (l *linkRepositoryImpl) GetFullLink(ctx context.Context, shortLinkTail string) (fullLink string, err error) {
	return l.client.Get(ctx, shortLinkTail).Result()
}
