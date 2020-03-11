
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	go build -o user-api *.go

.PHONY: run
run: build

	./user-api


.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-api:latest
