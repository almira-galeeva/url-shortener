package shortener

import "context"

type Service interface {
	GetShortUrl(ctx context.Context, originalUrl string) (string, error)
	GetOriginalUrl(ctx context.Context, shortlUrl string) (string, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}
