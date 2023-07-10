package shortener

import (
	"context"
	"errors"
	"testing"

	shortenerMocks "github.com/almira-galeeva/url-shortener/internal/repository/shortener/mocks"
	"github.com/almira-galeeva/url-shortener/internal/service/shortener"
	desc "github.com/almira-galeeva/url-shortener/pkg/shortener"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetShortUrl(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		shortUrl    = "https://shorturl.com/hjhjgjhlfgl"
		originalUrl = "https://github.com/almira-galeeva/url-shortener"
		invalidUrl  = "invalidUrl"
		repoErrText = gofakeit.Phrase()

		req = &desc.GetShortUrlRequest{
			OriginalUrl: originalUrl,
		}

		reqInvalid = &desc.GetShortUrlRequest{
			OriginalUrl: invalidUrl,
		}

		validRes = &desc.GetShortUrlResponse{
			ShortUrl: shortUrl,
		}

		repoErr = errors.New(repoErrText)
	)

	shortenerMock := shortenerMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		shortenerMock.EXPECT().CreateUrl(ctx, originalUrl, gomock.Any()).Return(false, nil),
		shortenerMock.EXPECT().CreateUrl(ctx, originalUrl, gomock.Any()).Return(false, repoErr),
	)

	api := newMockImplementation(Implementation{
		shortenerService: shortener.NewServiceMock(shortenerMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.GetShortUrl(ctx, req)

		require.Nil(t, err)
		require.NotEqual(t, validRes, res)
	})

	t.Run("repo err", func(t *testing.T) {
		_, err := api.GetShortUrl(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})

	t.Run("invalid url err", func(t *testing.T) {
		_, err := api.GetShortUrl(ctx, reqInvalid)
		require.Error(t, err)
	})
}
