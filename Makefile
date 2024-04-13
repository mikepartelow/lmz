all: build

test:
	@go test -cover ./...

build: test
	@go build -o . ./cmd/...

install: test
	@go install ./cmd/...
