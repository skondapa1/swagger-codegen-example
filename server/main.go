package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "petserver/go"
    dbapi "server/controller"
)

func main() {


    defer dbapi.DbClose()
    dbapi.DbInit()

	log.Printf("Server started")
	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
