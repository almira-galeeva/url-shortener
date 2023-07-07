LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=url-shortener user=url-shortener password=url_shortener-password sslmode=disable"

install-go-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	

PHONY: generate
generate:
	mkdir -p pkg/shortener_v1
	protoc --proto_path api/shortener_v1 --proto_path vendor.protogen \
		   --go_out=pkg/shortener_v1 --go_opt=paths=source_relative \
		   --plugin=protoc-gen-go=bin/protoc-gen-go \
		   --go-grpc_out=pkg/shortener_v1 --go-grpc_opt=paths=source_relative \
		   --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		   --grpc-gateway_out=pkg/shortener_v1 \
		   --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
		   --grpc-gateway_opt=logtostderr=true\
		   --grpc-gateway_opt=paths=source_relative \
		  api/shortener_v1/shortener.proto


PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/google/protobuf ]; then \
		git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
		mkdir -p  vendor.protogen/google/protobuf &&\
		mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
		rm -rf vendor.protogen/protobuf ;\
	fi

.PHONY: install-goose
.install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v