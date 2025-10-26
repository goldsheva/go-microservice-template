.PHONY: build
build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc go build -o bin/app-linux -v cmd/main.go
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o bin/app-macos -v cmd/main.go
