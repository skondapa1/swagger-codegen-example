#GOARCH=386
GOARCH=amd64
GOCMD=go
GOBUILD=GOARCH=${GOARCH} $(GOCMD) build
GOINSTALL=GOARCH=${GOARCH} $(GOCMD) install


all:
	swagger-codegen generate -i https://petstore.swagger.io/v2/swagger.json -l go-server  -o petserver
	swagger-codegen generate -i https://petstore.swagger.io/v2/swagger.json -l go  -o petstore
	cp ./server/override_patch/routers.go ./petserver/go/
	cp ./server/override_patch/pet_handler.go ./petserver/go/
	cp ./server/override_patch/pets_collection.go ./petserver/go/ 
	cp -r ./server/override_patch/controller ./petserver/go
	go get -d -v client/...
	go get -d -v server
	$(GOINSTALL) server
	$(GOINSTALL) client/... 

build: deps
	$(GOBUILD) server 
	$(GOBUILD) client/... 

clean:
	go clean -i server
	go clean -i client/...

deps: 
	go get -d -v client/...
	go get -d -v server

.DEFAULT_GOAL := all 

.PHONY: \
	all \
	build \
	clean \
	deps  \
	codegen-server  \
	codegen-client 



