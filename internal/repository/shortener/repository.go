package repository

import "context"

type Repository interface {
	GetShortUrl(ctx context.Context, originalUrl string) (string, error)
	GetOriginalUrl(ctx context.Context, shortlUrl string) (string, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetOriginalUrl(ctx context.Context, shortlUrl string) (string, error) {
	return "", nil
}

func (r *repository) GetShortUrl(ctx context.Context, originalUrl string) (string, error) {
	return "", nil
}
