all: test build

setup:
	@go get -u code.google.com/p/go.tools/cmd/vet
	@go get -u code.google.com/p/go.tools/cmd/cover
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/kisielk/errcheck


errcheck:
	@echo "=== errcheck ==="
	@errcheck ./...

vet:
	@echo "==== go vet ==="
	@go vet ./...

lint:
	@echo "==== go lint ==="
	@golint *.go

failfmt:
	@echo "===failfmt==="
	@script/failfmt

fmt:
	@echo "=== go fmt ==="
	@go fmt ./...

install: test
	@echo "=== go install ==="
	@go install -ldflags=$(GOLDFLAGS)

build:
	@echo "=== go build ==="
	@go build -ldflags=$(GOLDFLAGS) -o brigade

test: failfmt vet lint errcheck
	@echo "=== go test ==="
	@go test ./... -cover

.PHONY: setup cloc errcheck vet lint fmt install build test deploy
