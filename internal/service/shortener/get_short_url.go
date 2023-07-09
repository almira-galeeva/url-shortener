package shortener

import (
	"context"
	"math/rand"
	"net/url"
	"time"
)

func generateShortUrl(lenght int) (string, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	alphabet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_")

	shortUrl := make([]rune, lenght)

	for i := range shortUrl {
		shortUrl[i] = alphabet[rnd.Intn(len(alphabet))]
	}

	return string(shortUrl), nil
}

func (s *service) GetShortUrl(ctx context.Context, originalUrl string) (string, error) {
	_, err := url.ParseRequestURI(originalUrl)
	if err != nil {
		return "", err
	}

	shortUrl, err := generateShortUrl(s.urlLength)
	if err != nil {
		return "", err
	}

	// check if generated url has to be regenerated
	for {
		redoShort, err := s.shortenerRepository.CreateUrl(ctx, originalUrl, shortUrl)
		if err != nil {
			return "", err
		}

		if !redoShort {
			break
		}
	}

	return s.urlPrefix + shortUrl, nil
}
