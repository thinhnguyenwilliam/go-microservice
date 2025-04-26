APP_NAME = go-ecommerce-backend-api
MAIN_FILE = cmd/server/main.go

.PHONY: run build clean fmt lint test

run:
	go run $(MAIN_FILE)

build:
	go build -o $(APP_NAME) $(MAIN_FILE)

clean:
	go clean
	rm -f $(APP_NAME)

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...
