package repositories

import "context"

type LinkRepository interface {
	Create(ctx context.Context, linkToShorten string, shortLink string, ttlMinutes int) (err error)
	GetFullLink(ctx context.Context, shortLink string) (fullLink string, err error)
}
