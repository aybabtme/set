all: test

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

test: fmt vet lint errcheck
	@echo "=== go test ==="
	@go test ./... -v -cover

travistest: failfmt vet
	@echo "=== go test ==="
	@go test ./... -v -cover

.PHONY: setup cloc errcheck vet lint failfmt fmt test
