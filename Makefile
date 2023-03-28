version:=v0.1.0

all: clean dep compile

compile: build
	@echo "Done"

install:
	go install -ldflags="-X 'hyper-cli.Version=$(version)'" cmd/hyper.go

clean:
	@go clean && rm -rf bin

dep:
	@go mod tidy

build:
	@echo "Building"
	@go generate
	@go build -ldflags="-X 'hyper-cli.Version=$(version)'" -o 'bin/hyper' cmd/hyper.go

.PHONY: clean dep compile build
