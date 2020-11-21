#GOARCH=386
GOARCH=amd64
GOCMD=go
GOBUILD=GOARCH=${GOARCH} $(GOCMD) build
GOINSTALL=GOARCH=${GOARCH} $(GOCMD) install

codegen-server:
    swagger-codegen generate -i https://petstore.swagger.io/v2/swagger.json -l go-server  -o goserver

codegen-client:
    swagger-codegen generate -i https://petstore.swagger.io/v2/swagger.json -l go  -o petstore

all: codegen-server codegen-client deps
	$(GOINSTALL) server/... 
	$(GOINSTALL) client/... 

build: deps
	$(GOBUILD) server/... 
	$(GOBUILD) client/... 

clean:
	go clean -i server/...
	go clean -i client/...

deps:
	go get -d -v client/...
	go get -d -v server/...

.PHONY: \
	all \
	build \
	clean \
	deps  \
	codegen-server  \
	codegen-client 



