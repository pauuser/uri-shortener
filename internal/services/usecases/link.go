package usecases

import "context"

type LinkUseCase interface {
	Create(ctx context.Context, linkToShorten string, ttlMinutes int) (shortFullInk string, err error)
	GetFullLink(ctx context.Context, shortLink string) (fullLink string, err error)
}
