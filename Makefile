.PHONY: run test lint mock

run:
	go run cmd/app/main.go

test:
	go test ./...

lint:
	golangci-lint run --timeout=10m

mock:
	docker build --target mocks .

build:
	docker build -t owodunni/hano-scraper .
