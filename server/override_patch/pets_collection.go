package swagger

import (
   "reflect"
   "context"
   "fmt"
   "log"
   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo/options"
   db "petserver/go/controller"
)

var  pets  []Pet

func Find_ (slice interface{}, f func(value interface{}) bool) int {
    s := reflect.ValueOf(slice)
    if s.Kind() == reflect.Slice {
        for index := 0; index < s.Len(); index++ {
            if f(s.Index(index).Interface()) {
                return index
            }
        }
    }
    return -1
}

/**
 * InsertPet
 *
 * Insert a new pet object.
 */
func InsertPet_ (pet Pet) error {
    pets = append(pets, pet)
    return nil
}

func FindPetById_ (petId int64)  *Pet {
    idx :=  Find_(pets, func(value interface{}) bool {
       return value.(Pet).Id == petId
    })

    if (idx < 0) {
        return nil
    } else {
        fmt.Println("GetMyPet", pets[idx])
        return &pets[idx]
    }
}

/**
 * DeletePetById_
 *
 * Delete a pet by Id
 */
func DeletePetById_ (petId int64) int {

    idx :=  Find_(pets, func(value interface{}) bool {
       return value.(Pet).Id ==  petId
    })

    if (idx >= 0) {
        // changes order of pets
        pets[idx] = pets[len(pets)-1]
        pets = pets[:len(pets)-1]
    }

    return idx
}

func InsertPet (pet Pet) error {
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

func FindPetById (petId int64) *Pet {
    var pet Pet

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

func FindPetsByTag (tags []string) ([]Pet, error) {

    collection := db.CollectionPets()

    //tagsSlice := []bson.M{}
    var pets []Pet
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
            var elem Pet
            err := cur.Decode(&elem)
            if err != nil {
                log.Printf("%v\n", err)
                return nil, err
            }

            pets =append(pets, elem)
        }

        if err := cur.Err(); err != nil {
            log.Printf("%v\n", err)
            return nil, err
        }

        //Close the cursor once finished
        cur.Close(context.TODO())
        fmt.Printf("Found pets %v\n", pets)
    }
    return pets, nil
}
