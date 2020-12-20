package server

import (
   "context"
   "fmt"
   "log"
   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo/options"
   pets "swagger-codegen-example/pets/go"
   db "swagger-codegen-example/server/controller"
)

func InsertPet (pet pets.Pet) error {
    collection := db.CollectionPets()

    insertResult, err := 
          collection.InsertOne(context.TODO(), pet)

    if err != nil {
       log.Printf("Error inserting to Db : %v", err)
       return err
    }

    fmt.Println("Inserted post with ID:", insertResult.InsertedID)
    return nil
}

func FindPetById (petId int64) *pets.Pet {
    var pet pets.Pet

    collection := db.CollectionPets()
    filter := bson.M { "id" : petId}

    err := collection.FindOne(context.TODO(), filter).Decode(&pet)
    if err != nil {
       log.Printf("Pet id %d Not found in Db : %v", petId, err)
       return nil
    }

    fmt.Printf("Found pet with Name %s\n", pet.Name)
    return &pet
}

func DeletePetById (petId int64) int {
    collection := db.CollectionPets()
    filter := bson.M { "id" : petId}

    result, err := 
        collection.DeleteOne(context.TODO(), filter)

    if err != nil {
       log.Printf("Pet id %d Not found in Db : %v", petId, err)
       return 0
    }

    fmt.Printf("DeleteOne removed %v document(s)\n", 
            result.DeletedCount)

    return int(result.DeletedCount)
}

/*
   var pets []bson.M
   if err = cursor.All(ctx, &pets); err != nil {
       log.Fatal(err)
   }
*/

func FindPetsByTag (tags []string) ([]pets.Pet, error) {

    collection := db.CollectionPets()

    //tagsSlice := []bson.M{}
    var petsArr []pets.Pet
    for _, e  := range tags {
        //tagsSlice = append(tagsSlice, 
              //bson.M{"$elemMatch": bson.M{"name": e}})
        tagsSlice := bson.M{"$elemMatch": bson.M{"name": e}}

        //filter := bson.M { 
             //"tags" :  bson.M { "all": tagsSlice},
        //}

        filter := bson.M { 
             "tags" :  tagsSlice,
        }


    //{"$all": tags_slice},
    //d  := bson.D { 
           //{ Name: "name", Value: "lab" },
           //{ Name: "id", Value: 1 },
        //}


        //find records
        //pass these options to the Find method
        findOptions := options.Find()

        //Set the limit of the number of record to find
        findOptions.SetLimit(5)

        //Define an array in which you can store the decoded documents
        fmt.Println("filter %v", filter)

        cur, err := collection.Find(context.TODO(), filter, findOptions)
        if err != nil {
           log.Printf("Pets matching tags %v Not found in Db : %v", 
                  tags, err)
           return nil, err
        }


        //Finding multiple documents returns a cursor
        //Iterate through the cursor allows us to decode documents one at a time
        for cur.Next(context.TODO()) {
            //Create a value into which the single document can be decoded
            var elem pets.Pet
            err := cur.Decode(&elem)
            if err != nil {
                log.Printf("%v\n", err)
                return nil, err
            }

            petsArr =append(petsArr, elem)
        }

        if err := cur.Err(); err != nil {
            log.Printf("%v\n", err)
            return nil, err
        }

        //Close the cursor once finished
        cur.Close(context.TODO())
        fmt.Printf("Found pets %v\n", petsArr)
    }
    return petsArr, nil
}
