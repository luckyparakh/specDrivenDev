.PHONY: build dev lint dev-air dev-templ dev-css migrate sqlc

build:
	templ generate
	go build -o bin/agentclinic$(if $(filter Windows_NT,$(OS)),.exe,) .

dev:
	$(MAKE) -j3 dev-templ dev-css dev-air

dev-air:
	air

dev-templ:
	templ generate --watch

dev-css:
	npm run css -- --watch

lint:
	go vet ./...

migrate:
	go run ./cmd/migrate up

sqlc:
	sqlc generate
