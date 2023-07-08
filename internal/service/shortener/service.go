package shortener

import (
	"context"

	shortenerRepository "github.com/almira-galeeva/url-shortener/internal/repository/shortener"
)

var urlPrefix = "https://shorturl.com/"

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
