package controller

import (
   "context"
   "fmt"
   "time"
   "log"
   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
   "go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
    // Name of the database.
    DBName = "my_database"
    // URI = "mongodb://<user>:<password>@<host>/<name>"
    URI = "mongodb://localhost:27017"
)

var db  *mongo.Database
var client  *mongo.Client

var bgCtx = context.Background()

func setupClient(opts ...*options.ClientOptions) *mongo.Client {
	if len(opts) == 0 {
		opts = append(opts, options.Client().ApplyURI(URI))
	}
	client, _ := mongo.NewClient(opts...)
	return client
}


func DbInit() error {
    var err error

    client =  setupClient()
    ctx, _ := context.WithTimeout(bgCtx, 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
       log.Print(err)
       return err
    }

    // Ping our db connection
    err = client.Ping(bgCtx, readpref.Primary())
    if err != nil {
        log.Print("Couldn't connect to the database", err)
        return err
    } else {
        log.Println("Connected!")
    }


    db = client.Database(DBName)
    fmt.Println(db.Name()) // output: my_database

    collection := db.Collection("pets")
    /* 
     * D objects are actually just arrays of ordered E structs, 
     * which represent individual key-value pair elements. 
     * Finally, A is for (obviously) arrays of interface{} values.
     */
    if collection != nil {
       collection.Drop(ctx)
    }

    db.RunCommand(context.TODO(), bson.D{{"create", "pets"}})

    return nil
}

func DbClose() {
    ctx, _ := context.WithTimeout(bgCtx, 10*time.Second)
    client.Disconnect(ctx)
}

func CollectionPets() *mongo.Collection {
    return db.Collection("pets")
}


