package shortener

import (
	shortenerService "github.com/almira-galeeva/url-shortener/internal/service/shortener"
	desc "github.com/almira-galeeva/url-shortener/pkg/shortener"
)

type Implementation struct {
	desc.UnimplementedShortenerServer

	shortenerService shortenerService.Service
}

func NewImplementation(shortenerService shortenerService.Service) *Implementation {
	return &Implementation{
		shortenerService: shortenerService,
	}
}

func newMockImplementation(i Implementation) *Implementation {
	return &Implementation{
		desc.UnimplementedShortenerServer{},
		i.shortenerService,
	}
}
