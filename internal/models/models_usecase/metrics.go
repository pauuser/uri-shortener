package models_usecase

import "time"

type LinkMetrics struct {
	ClickCount   int64
	CreatedAtUtc time.Time
}
