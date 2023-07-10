package shortener

import (
	"context"
)

func (s *service) GetOriginalUrl(ctx context.Context, shortUrl string) (string, error) {
	originUrl, err := s.shortenerRepository.GetOriginalUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}

	return originUrl, nil
}
