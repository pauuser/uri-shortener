package repositories

import "context"

type LinkRepository interface {
	Create(ctx context.Context, linkToShorten string, shortLinkTail string, ttlMinutes int) (err error)
	GetFullLink(ctx context.Context, shortLinkTail string) (fullLink string, err error)
}
