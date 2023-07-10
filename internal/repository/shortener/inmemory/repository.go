package inmemory

import (
	"context"
	"fmt"
	"sync"

	irepo "github.com/almira-galeeva/url-shortener/internal/repository/shortener"
)

type repository struct {
	shortUrls    map[string]string
	originalUrls map[string]struct{}
	m            sync.RWMutex
}

func NewRepository() irepo.Repository {
	return &repository{
		shortUrls:    make(map[string]string),
		originalUrls: make(map[string]struct{}),
	}
}

func (r *repository) CreateUrl(ctx context.Context, originalUrl string, shortUrl string) (bool, error) {
	// check if original url already is in memory
	r.m.RLock()
	if _, ok := r.originalUrls[originalUrl]; ok {
		return false, fmt.Errorf("Url %s already exists in memory", originalUrl)
	}
	r.m.RUnlock()

	// check if generated url has to be regenerated
	r.m.RLock()
	if _, ok := r.shortUrls[shortUrl]; ok {
		return true, nil
	}
	r.m.RUnlock()

	// insert originalUrl, shortUrl in memory
	r.m.Lock()
	r.shortUrls[shortUrl] = originalUrl
	r.m.Unlock()

	// add originalUrl in hashset
	r.m.Lock()
	r.originalUrls[originalUrl] = struct{}{}
	r.m.Unlock()

	return false, nil
}

func (r *repository) GetOriginalUrl(ctx context.Context, shortlUrl string) (string, error) {
	r.m.RLock()
	defer r.m.RUnlock()
	if originalUrl, ok := r.shortUrls[shortlUrl]; ok {
		return originalUrl, nil
	}

	return "", fmt.Errorf("Original url by %s not found", shortlUrl)
}
