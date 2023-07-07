package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	impl "github.com/almira-galeeva/url-shortener/internal/api/shortener"
	config "github.com/almira-galeeva/url-shortener/internal/config"
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

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := runHTTP(ctx, cfg)
		if err != nil {
			log.Fatalf("Failed to run HTTP server: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		err := runGRPC(ctx, cfg)
		if err != nil {
			log.Fatalf("Failed to run gRPC server: %s", err.Error())
		}
	}()

	wg.Wait()
}

func runHTTP(ctx context.Context, cfg *config.Config) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := desc.RegisterShortenerHandlerFromEndpoint(ctx, mux, cfg.GRPC.GetAddress(), opts)
	if err != nil {
		return err
	}

	log.Printf("HTTP Server is running on host: %s", cfg.HTTP.GetAddress())

	return http.ListenAndServe(cfg.HTTP.GetAddress(), mux)
}

func runGRPC(ctx context.Context, cfg *config.Config) error {
	listener, err := net.Listen("tcp", cfg.GRPC.GetAddress())
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pgCfg, err := pgxpool.ParseConfig(cfg.DB.DSN)
	if err != nil {
		return err
	}

	dbc, err := pgxpool.ConnectConfig(ctx, pgCfg)
	if err != nil {
		return err
	}

	var shortenerRepository iShortenerRepo.Repository
	if cfg.DB.Source == "inmemory" {
		shortenerRepository = inMemoryShortenerRepo.NewRepository()
	} else if cfg.DB.Source == "db" {
		shortenerRepository = dbShortenerRepo.NewRepository(dbc)
	}

	shortenerService := shortenerService.NewService(shortenerRepository)
	desc.RegisterShortenerServer(s, impl.NewImplementation(shortenerService))

	log.Printf("GRPC Server is running on host: %s", cfg.GRPC.GetAddress())
	if err = s.Serve(listener); err != nil {
		return err
	}

	return nil
}
