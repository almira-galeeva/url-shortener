package shortener

import "context"

func (s *service) GetShortUrl(ctx context.Context, originalUrl string) (string, error) {
	shortUrl := "hjgfhjl"
	err := s.shortenerRepository.CreateUrl(ctx, originalUrl, shortUrl)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}
