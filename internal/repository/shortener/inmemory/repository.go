package inmemory

import (
	"context"
	"fmt"
	"sync"

	irepo "github.com/almira-galeeva/url-shortener/internal/repository/shortener"
)

const tableName = "urls"

type repository struct {
	urls map[string]string
	m    sync.RWMutex
}

func NewRepository() irepo.Repository {
	return &repository{
		urls: make(map[string]string),
	}
}

func (r *repository) CreateUrl(ctx context.Context, originalUrl string, shortUrl string) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.urls[shortUrl] = originalUrl
	return nil
}

func (r *repository) GetOriginalUrl(ctx context.Context, shortlUrl string) (string, error) {
	r.m.RLock()
	defer r.m.RUnlock()
	if originalUrl, ok := r.urls[shortlUrl]; ok {
		return originalUrl, nil
	}

	return "", fmt.Errorf("Original url by %s not found", shortlUrl)
}
