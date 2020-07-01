.PHONY: build
build:
		go build -o balance.exe -v ./cmd/server

.DEFAULT_GOAL := build

.PHONY: migrate_up
migrate_up:
		migrate -path migrations/ -database "postgres://postgres:postgres@localhost/balance?sslmode=disable" up

.PHONY: migrate_down
migrate_down:
		migrate -path migrations/ -database "postgres://postgres:postgres@localhost/balance?sslmode=disable" down

.PHONY: migrate_force
migrate_force:
		migrate -path migrations/ -database "postgres://postgres:postgres@localhost/balance?sslmode=disable" force 20200629183020

.PHONY: migrate_test_up
migrate_test_up:
		migrate -path migrations/ -database "postgres://postgres:postgres@localhost/balance_test?sslmode=disable" up

.PHONY: migrate_test_down
migrate_test_down:
		migrate -path migrations/ -database "postgres://postgres:postgres@localhost/balance_test?sslmode=disable" down

.PHONY: migrate_test_force
migrate_test_force:
		migrate -path migrations/ -database "postgres://postgres:postgres@localhost/balance_test?sslmode=disable" force 20200629183020

