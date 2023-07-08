package shortener

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

func (s *service) GetOriginalUrl(ctx context.Context, shortUrl string) (string, error) {
	_, err := url.ParseRequestURI(shortUrl)
	if err != nil {
		return "", err
	}

	urlPrefix := "https://shorturl/"
	if !strings.HasPrefix(shortUrl, urlPrefix) {
		return "", fmt.Errorf("Short url should start with %s", urlPrefix)
	}

	originUrl, err := s.shortenerRepository.GetOriginalUrl(ctx, shortUrl[len(urlPrefix):])
	if err != nil {
		return "", err
	}

	return originUrl, nil
}
