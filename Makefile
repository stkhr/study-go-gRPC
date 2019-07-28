# Go パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=handle
BINARY_UNIX=$(BINARY_NAME)_unix

download:
	$(GOMOD) download

build:
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/lambda/spot-interruption/main.go

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v cmd/lambda/spot-interruption/main.go

test:
	$(GOTEST) -v ./...