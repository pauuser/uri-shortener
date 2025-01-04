package usecases

import (
	"context"
	"uri-shortener/internal/models/models_usecase"
)

type LinkUseCase interface {
	Create(ctx context.Context, linkToShorten string, ttlMinutes int) (shortFullInk string, err error)
	GetFullLink(ctx context.Context, shortLink string) (fullLink string, err error)
	GetMetrics(ctx context.Context, shortLink string) (metrics models_usecase.LinkMetrics, err error)
}
