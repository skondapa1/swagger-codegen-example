#GOARCH=386
#prefix=$(shell bash -c pwd)
#export GOPATH=${prefix}
GOARCH=amd64
GOCMD=go
GOBUILD=GOARCH=${GOARCH} $(GOCMD) build
GOINSTALL=GOARCH=${GOARCH} $(GOCMD) install


#generate model objects from spec file
#generate server source without the model
#generate client source

generate: 
	swagger-codegen generate -i https://petstore.swagger.io/v2/swagger.json -l go  -o petstore -c config_client.json
	swagger-codegen generate -i https://petstore.swagger.io/v2/swagger.json -l go-server -o pets -t ./resources/ -Dmodels -c config.json 
	swagger-codegen generate -i https://petstore.swagger.io/v2/swagger.json -l go-server  -o petstore_server -t ./resources/ --import-mappings Pet=swagger-codegen-example/pets,Order=swagger-codegen-example/pets,Body1=swagger-codegen-example/pets,ApiResponse=swagger-codegen-example/pets,Category=swagger-codegen-example/pets,Body=swagger-codegen-example/pets,Order=swagger-codegen-example/pets,Tag=swagger-codegen-example/pets,User=swagger-codegen-example/pets -c config_server.json

deps:
	go get -d -v swagger-codegen-example/...

build: deps
	$(GOINSTALL) swagger-codegen-example/petstore_server/...
	$(GOINSTALL) swagger-codegen-example/client/... 

all: build

clean:
	-rm -rf petstore_server
	-rm -rf petstore
	-rm -rf pets
	-go clean -i swagger-codegen-example/client/...
	-go clean -i swagger-codegen-example/server/...
	-rm -f $(GOPATH)/bin/client
	-rm -f $(GOPATH)/bin/petstore_server


.DEFAULT_GOAL := all 

.PHONY:  all build clean  deps  



