package repositories

import (
	"context"
	"uri-shortener/internal/models/models_usecase"
)

type LinkRepository interface {
	Create(ctx context.Context, linkToShorten string, shortLinkTail string, ttlMinutes int) (err error)
	GetFullLink(ctx context.Context, shortLinkTail string, incrementClickCount bool) (fullLinkModel *models_usecase.Link, err error)
}
