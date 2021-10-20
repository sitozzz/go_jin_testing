all: run

run:
	@go run ./main.go

build:
	@go build -v .

lint:
	@golint ./...

help:
	@echo "make: run in dev mode"
	@echo "make build: build binary executable"
	@echo "make lint: run golint to all project"
