package models_repo

import (
	"time"
	"uri-shortener/internal/models/models_usecase"
)

type LinkDb struct {
	FullLink     string    `json:"full_link"`
	ClickCount   int64     `json:"click_count"`
	CreatedAtUtc time.Time `json:"created_at_utc"`
}

func (l *LinkDb) ToUsecase() *models_usecase.Link {
	return &models_usecase.Link{
		FullLink:     l.FullLink,
		ClickCount:   l.ClickCount,
		CreatedAtUtc: l.CreatedAtUtc,
	}
}
