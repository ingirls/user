
GOPATH:=$(shell go env GOPATH)


.PHONY: proto-srv
proto-srv:
    
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto

.PHONY: proto-api
proto-api:
    
	cd api; protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto
    

.PHONY: build-srv
build-srv: 
	go build -mod=vendor -o user-srv *.go

.PHONY: build-api
build-api: 
	cd api; go build -mod=vendor -o user-api *.go


.PHONY: srv
srv: build-srv
	./user-srv --registry=consul --registry_address=consul:8500

.PHONY: api
api: build-api
	./api/user-api --registry=consul --registry_address=consul:8500


.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-srv:latest
