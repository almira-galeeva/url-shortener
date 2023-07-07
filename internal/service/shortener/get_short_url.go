package shortener

import "context"

func (s *service) GetShortUrl(ctx context.Context, originalUrl string) (string, error) {
	shortUrl, err := s.shortenerRepository.GetShortUrl(ctx, originalUrl)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}
