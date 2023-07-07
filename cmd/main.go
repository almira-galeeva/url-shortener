package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	impl "github.com/almira-galeeva/url-shortener/internal/api/shortener"
	shortenerRepository "github.com/almira-galeeva/url-shortener/internal/repository/shortener"
	shortenerService "github.com/almira-galeeva/url-shortener/internal/service/shortener"
	desc "github.com/almira-galeeva/url-shortener/pkg/shortener"
)

const (
	grpcHost = "localhost:50051"
	httpHost = "localhost:8080"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := runHTTP()
		if err != nil {
			log.Fatalf("Failed to run HTTP server: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		err := runGRPC()
		if err != nil {
			log.Fatalf("Failed to run gRPC server: %s", err.Error())
		}
	}()

	wg.Wait()
}

func runHTTP() error {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := desc.RegisterShortenerHandlerFromEndpoint(ctx, mux, grpcHost, opts)
	if err != nil {
		return err
	}

	log.Printf("HTTP Server is running on host: %s", httpHost)

	return http.ListenAndServe(httpHost, mux)
}

func runGRPC() error {
	listener, err := net.Listen("tcp", grpcHost)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)

	shortenerRepository := shortenerRepository.NewRepository()
	shortenerService := shortenerService.NewService(shortenerRepository)
	desc.RegisterShortenerServer(s, impl.NewImplementation(shortenerService))

	log.Printf("GRPC Server is running on host: %s", grpcHost)
	if err = s.Serve(listener); err != nil {
		return err
	}

	return nil
}
