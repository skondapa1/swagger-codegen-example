# SWAGGER-CODEGEN example 
# REFERENCES
   <br> https://levelup.gitconnected.com/tools-for-implementing-a-golang-api-server-with-auto-generated-code-and-documentation-694262e3866c
   <br> https://goswagger.io/tutorial/todo-list.html
   <br> https://github.com/deepmap/oapi-codegen
   <br> https://github.com/OAI/OpenAPI-Specification
   <br> https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/
   <br> https://www.mongodb.com/golang
   <br> https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/
   <br> This page describes the http request structure -
   <br> https://golang.org/pkg/net/http/#Request
   <br> REST-API with Golang and Mux
   <br> https://medium.com/@hugo.bjarred/rest-api-with-golang-and-mux-e934f581b8b5
   <br> https://medium.com/@hugo.bjarred/rest-api-with-golang-mux-mysql-c5915347fa5b
   <br> https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver
 

<b> -installing swagger-codegen </b>
<br> brew install swagger-codegen

<b> -installing mongodb on ubuntu and starting </b>
<br> wget -qO - https://www.mongodb.org/static/pgp/server-4.4.asc | sudo apt-key add -
<br> touch /etc/apt/sources.list.d/mongodb-org-4.4.list
<br> echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.4.list
<br> sudo apt-get update
<br> sudo apt-get install -y mongodb-org
<br> sudo systemctl daemon-reload
<br> sudo systemctl start mongod
<br> sudo systemctl status mongod

<b> -client code generate </b>
<br> swagger-codegen generate -i https://petstore.swagger.io/v2/swagger.json -l go  -o ./go/src/petstore

<b> -server code generate </b>
<br> swagger-codegen generate -i https://petstore.swagger.io/v2/swagger.json -l go-server  -o ./go/src/petserver

<b> -build your client and server, setup GOPATH correctly if not set. </b>
<b> -update dependencies, update mongodb go driver. </b>
<br> go get go.mongodb.org/mongo-driver
<br> go get client
<br> go get server

<br> go install client
<br> go install server
 
