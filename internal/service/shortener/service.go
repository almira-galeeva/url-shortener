package shortener

import (
	"context"

	shortenerRepository "github.com/almira-galeeva/url-shortener/internal/repository/shortener"
)

type Service interface {
	GetShortUrl(ctx context.Context, originalUrl string) (string, error)
	GetOriginalUrl(ctx context.Context, shortUrl string) (string, error)
}

type service struct {
	shortenerRepository shortenerRepository.Repository
}

func NewService(shortenerRepository shortenerRepository.Repository) Service {
	return &service{
		shortenerRepository: shortenerRepository,
	}
}
