#!/bin/bash

MIGRATION_DIR=./migrations
export MIGRATION_DSN="host=db port=5432 dbname=shortener-service user=shortener-service-user password=shortener-password sslmode=disable"

sleep 2 && goose -dir ${MIGRATION_DIR} postgres "${MIGRATION_DSN}" up -v