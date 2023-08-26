.PHONY: deps
deps:
	@go mod download
	@go mod verify

.PHONY: migrate_up
migrate_up:
	migrate -database "postgres://postgres:1234@db:5432/postgres?sslmode=disable" -path migrations up

.PHONY: wait
wait:
	@chmod +x ./wait-for-postgres.sh
	@./wait-for-postgres.sh db

.PHONY: build
build: deps
	@go build -buildvcs=false -o /usr/local/bin/server ./...


.PHONY: run
run: wait migrate_up build
	@server