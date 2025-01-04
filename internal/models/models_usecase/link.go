package models_usecase

import "time"

type Link struct {
	FullLink     string
	ClickCount   int64
	CreatedAtUtc time.Time
}
