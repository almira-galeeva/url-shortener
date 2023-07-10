LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=shortener-service user=shortener-service-user password=shortener-password sslmode=disable"

install-go-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	

generate:
	mkdir -p pkg/shortener
	protoc --proto_path api/shortener \
		   --go_out=pkg/shortener --go_opt=paths=source_relative \
		   --plugin=protoc-gen-go=bin/protoc-gen-go \
		   --go-grpc_out=pkg/shortener --go-grpc_opt=paths=source_relative \
		   --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		  api/shortener/shortener.proto

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

run:
	docker-compose up -d
	sleep 10
	go build -o shortener-service cmd/main.go
	./shortener-service