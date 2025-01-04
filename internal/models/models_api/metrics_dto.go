package models_api

import (
	"time"
	"uri-shortener/internal/models/models_usecase"
)

type LinkMetricsDto struct {
	ClickCount   int64     `json:"clickCount"`
	CreatedAtUtc time.Time `json:"created_at_utc"`
}

func ToDto(l models_usecase.LinkMetrics) LinkMetricsDto {
	return LinkMetricsDto{
		ClickCount:   l.ClickCount,
		CreatedAtUtc: l.CreatedAtUtc,
	}
}
