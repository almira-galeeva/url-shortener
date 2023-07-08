package shortener

import (
	"context"
	"strings"
)

func generateShortUrl(originalUrl string) (string, error) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
	shortUrl := make([]string, 0, 10)
	shortUrl = append(shortUrl, "https://shorturl/")

	for _, byte := range originalUrl {
		ind := int(byte) % len(alphabet)
		shortUrl = append(shortUrl, string(alphabet[ind]))
	}

	return strings.Join(shortUrl, ""), nil
}

func (s *service) GetShortUrl(ctx context.Context, originalUrl string) (string, error) {
	shortUrl, err := generateShortUrl(originalUrl)
	if err != nil {
		return "", err
	}

	err = s.shortenerRepository.CreateUrl(ctx, originalUrl, shortUrl)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}
