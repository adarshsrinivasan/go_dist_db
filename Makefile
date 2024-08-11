refresh-vendor:
	go mod tidy
	go mod download
	go mod vendor

build:
	@go vet ./...
	@go fmt ./...
	@go build -o bin/fs

run: build
	@./bin/fs

test:
	go test ./... -v