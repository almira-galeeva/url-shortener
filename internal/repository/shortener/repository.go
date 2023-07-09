package repository

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_repository.go -package=mocks . Repository

import (
	"context"
)

type Repository interface {
	CreateUrl(ctx context.Context, originalUrl string, shortUrl string) (bool, error)
	GetOriginalUrl(ctx context.Context, shortlUrl string) (string, error)
}
