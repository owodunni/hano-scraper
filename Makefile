.PHONY: run test lint mock

run:
	go run cmd/app/main.go

test:
	go test ./...

lint:
	golangci-lint run --timeout=10m

mock:
	go generate -run "mockgen" ./...	

build:
	docker build -t owodunni/hano-scraper .
