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
	urlPrefix           string
	urlLength           int
}

func NewService(shortenerRepository shortenerRepository.Repository, urlPrefix string, urlLength int) Service {
	return &service{
		shortenerRepository: shortenerRepository,
		urlPrefix:           urlPrefix,
		urlLength:           urlLength,
	}
}

func NewServiceMock(deps ...interface{}) Service {
	is := service{}

	for _, v := range deps {
		switch s := v.(type) {
		case shortenerRepository.Repository:
			is.shortenerRepository = s
		}
	}
	return &is
}
