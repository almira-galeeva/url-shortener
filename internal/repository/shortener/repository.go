package repository

import (
	"context"
)

type Repository interface {
	CreateUrl(ctx context.Context, originalUrl string, shortUrl string) error
	GetOriginalUrl(ctx context.Context, shortlUrl string) (string, error)
}
