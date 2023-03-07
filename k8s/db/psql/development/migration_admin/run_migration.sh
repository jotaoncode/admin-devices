#!/usr/bin/env sh

user=$(tail -1 /etc/secret/POSTGRES_USER)
pass=$(tail -1 /etc/secret/POSTGRES_PASSWORD)

migrate -database postgres://$user:$pass@$POSTGRES_SERVICE_HOST:$POSTGRES_SERVICE_PORT/admin?sslmode=disable -path ./sql up