package repositories_impl

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
	"uri-shortener/internal/models/models_repo"
	"uri-shortener/internal/models/models_usecase"
	"uri-shortener/internal/pkg/errors/usecase_errors"
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
	link := models_repo.LinkDb{
		FullLink:     linkToShorten,
		CreatedAtUtc: time.Now().UTC(),
		ClickCount:   0,
	}
	linkJson, err := json.Marshal(link)
	if err != nil {
		return err
	}
	err = l.client.Set(ctx, shortLinkTail, string(linkJson), expiration).Err()

	return err
}

func (l *linkRepositoryImpl) GetFullLink(ctx context.Context, shortLinkTail string, incrementClickCount bool) (fullLinkModel *models_usecase.Link, err error) {
	fullLink, err := l.client.Get(ctx, shortLinkTail).Result()
	fullLinkJson := models_repo.LinkDb{}
	if errors.Is(err, redis.Nil) {
		return nil, usecase_errors.LinkNotFoundError
	}

	err = json.Unmarshal([]byte(fullLink), &fullLinkJson)
	if incrementClickCount {
		fullLinkJson.ClickCount++
		linkJson, err := json.Marshal(fullLinkJson)
		if err != nil {
			return nil, err
		}
		err = l.client.Set(ctx, shortLinkTail, string(linkJson), redis.KeepTTL).Err()
	}

	return fullLinkJson.ToUsecase(), err
}
