package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	httpImpl "github.com/almira-galeeva/url-shortener/internal/api/http"
	grpcImpl "github.com/almira-galeeva/url-shortener/internal/api/shortener"
	"github.com/almira-galeeva/url-shortener/internal/config"
	iShortenerRepo "github.com/almira-galeeva/url-shortener/internal/repository/shortener"
	dbShortenerRepo "github.com/almira-galeeva/url-shortener/internal/repository/shortener/db"
	inMemoryShortenerRepo "github.com/almira-galeeva/url-shortener/internal/repository/shortener/inmemory"
	shortenerService "github.com/almira-galeeva/url-shortener/internal/service/shortener"
	desc "github.com/almira-galeeva/url-shortener/pkg/shortener"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "config/config.json", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	cfg, err := config.NewConfig(pathConfig)
	if err != nil {
		log.Fatalf("Failed to parse config: %s", err.Error())
	}

	pgCfg, err := pgxpool.ParseConfig(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("Failed to parse DSN: %s", err.Error())
	}

	dbc, err := pgxpool.ConnectConfig(ctx, pgCfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s", err.Error())
	}

	var shortenerRepository iShortenerRepo.Repository
	if cfg.DB.Source == "inmemory" {
		shortenerRepository = inMemoryShortenerRepo.NewRepository()
	} else if cfg.DB.Source == "db" {
		shortenerRepository = dbShortenerRepo.NewRepository(dbc)
	}

	shortenerSrv := shortenerService.NewService(shortenerRepository, cfg.UrlPrefix, cfg.UrlLength)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err = runHTTP(shortenerSrv, cfg.HTTP)
		if err != nil {
			log.Fatalf("Failed to run HTTP server: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		err = runGRPC(shortenerSrv, cfg.GRPC)
		if err != nil {
			log.Fatalf("Failed to run gRPC server: %s", err.Error())
		}
	}()

	wg.Wait()
}

func runHTTP(shortenerSrv shortenerService.Service, cfg config.HTTP) error {
	router := mux.NewRouter()
	router.StrictSlash(true)

	impl := httpImpl.NewImplementation(shortenerSrv)

	router.HandleFunc("/shortener/short_url", impl.GetShortUrl).Methods("POST")
	router.HandleFunc("/shortener/original_url/{short_url}", impl.GetOriginalUrl).Methods("GET")

	log.Printf("HTTP Server is running on host: %s", cfg.GetAddress())
	return http.ListenAndServe(cfg.GetAddress(), router)

}

func runGRPC(shortenerSrv shortenerService.Service, cfg config.GRPC) error {
	listener, err := net.Listen("tcp", cfg.GetAddress())
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterShortenerServer(s, grpcImpl.NewImplementation(shortenerSrv))

	log.Printf("GRPC Server is running on host: %s", cfg.GetAddress())
	if err = s.Serve(listener); err != nil {
		return err
	}

	return nil
}
