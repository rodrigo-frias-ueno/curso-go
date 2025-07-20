.PHONY: build run download tests mocks swagger

init: download mocks swagger

build:
	go build -o app -v ./cmd/restapi/

run:
	go run cmd/restapi/*

download:
	go mod download

tests:
	go test ./...

mocks: 
	go install go.uber.org/mock/mockgen@latest
	go generate ./...
	go mod tidy

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -d cmd/restapi,internal/rest --parseInternal
	go mod tidy