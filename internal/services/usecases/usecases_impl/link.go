package usecases_impl

import (
	"context"
	"math/rand"
	"uri-shortener/cmd/modes/flags"
	"uri-shortener/internal/services/repositories"
	"uri-shortener/internal/services/usecases"
)

type linkUseCase struct {
	repo      repositories.LinkRepository
	linkFlags flags.LinkFlags
}

func (l linkUseCase) Create(ctx context.Context, linkToShorten string, ttlMinutes int) (shortFullLink string, err error) {
	tail := l.GenerateRandomTail(l.linkFlags.LinkLen)
	shortLink := l.linkFlags.Domain + "/" + tail

	err = l.repo.Create(ctx, linkToShorten, tail, ttlMinutes)

	return shortLink, err
}

func (l linkUseCase) GenerateRandomTail(tailLen int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

	b := make([]byte, tailLen)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func (l linkUseCase) GetFullLink(ctx context.Context, shortLinkTail string) (fullLink string, err error) {
	return l.repo.GetFullLink(ctx, shortLinkTail)
}

func NewLinkUseCase(repository repositories.LinkRepository, configuration flags.LinkFlags) usecases.LinkUseCase {
	return &linkUseCase{
		repo:      repository,
		linkFlags: configuration,
	}
}
