package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	shortenerService "github.com/almira-galeeva/url-shortener/internal/service/shortener"
)

type Implementation struct {
	shortenerService shortenerService.Service
}

func NewImplementation(shortenerService shortenerService.Service) *Implementation {
	return &Implementation{
		shortenerService: shortenerService,
	}
}

func (i *Implementation) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	type shortUrlRequest struct {
		OriginalUrl string `json:"original_url"`
	}

	request := shortUrlRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortUrl, err := i.shortenerService.GetShortUrl(r.Context(), request.OriginalUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(shortUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (i *Implementation) GetOriginalUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["short_url"]

	originalUrl, err := i.shortenerService.GetOriginalUrl(r.Context(), shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(originalUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
