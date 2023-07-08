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

func TestGetOriginalUrl(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		shortUrl    = "https://shorturl.com/hjhjgjhlfgl"
		originalUrl = gofakeit.URL()
		repoErrText = gofakeit.Phrase()

		req = &desc.GetOriginalUrlRequest{
			ShortUrl: shortUrl,
		}

		validRes = &desc.GetOriginalUrlResponse{
			OriginalUrl: originalUrl,
		}

		repoErr = errors.New("Error")
	)

	shortenerMock := shortenerMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		shortenerMock.EXPECT().GetOriginalUrl(ctx, req).Return(originalUrl, nil),
		shortenerMock.EXPECT().GetOriginalUrl(ctx, req).Return("", repoErr),
	)

	api := newMockImplementation(Implementation{
		shortenerService: shortener.NewServiceMock(shortenerMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.GetOriginalUrl(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("repo err", func(t *testing.T) {
		_, err := api.GetOriginalUrl(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
