package shortener

import (
	"context"

	desc "github.com/almira-galeeva/url-shortener/pkg/shortener"
)

func (i *Implementation) GetOriginalUrl(ctx context.Context, req *desc.GetOriginalUrlRequest) (*desc.GetOriginalUrlResponse, error) {
	originalUrl, err := i.shortenerService.GetOriginalUrl(ctx, req.GetShortUrl())
	if err != nil {
		return nil, err
	}

	return &desc.GetOriginalUrlResponse{
		OriginalUrl: originalUrl,
	}, nil
}
