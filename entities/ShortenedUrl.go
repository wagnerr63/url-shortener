package entities

import (
	"time"

	"github.com/google/uuid"
)

type ShortenedUrl struct {
	ID          string `gorm:"primary_key" json:"id"`
	Url         string `gorm:"url" json:"url"`
	ShortenedID string `gorm:"shortened_id" json:"shortened_id"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	ExpiresIn   string `db:"expires_in" json:"expires_in"`
}

func NewShortenedUrl() ShortenedUrl {
	return ShortenedUrl{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		ExpiresIn: time.Now().AddDate(0, 0, 1).Format("2006-01-02 15:04:05"),
	}
}
