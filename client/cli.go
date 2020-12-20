package main

//import "golang.org/x/oauth2"
import "swagger-codegen-example/petstore"
import "fmt"
import "golang.org/x/net/context"



func  main() {

    cfg := &petstore.Configuration{
		BasePath:      "http://localhost:8080/v2",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}

    category := &petstore.Category {
      1,
      "Dog",
    }

    tags := []petstore.Tag {
       petstore.Tag{1, "lab"},
       petstore.Tag{2, "labradoodle"}, 
       petstore.Tag{3, "GermanSheperd"},
       petstore.Tag{4, "AustralianDoberman"},
     }

    rascal := petstore.Pet {
      1234,
      category,
      "rascal",
      []string{},
      tags[0:1],
      "7 yr old dog",
    }

    coffee := petstore.Pet {
      1235,
      category,
      "coffee",
      []string{},
      tags[0:2],
      "7 yr old dog",
    }

    client :=  petstore.NewAPIClient(cfg)

    add_r, add_err := client.PetApi.AddPet(context.Background(), rascal) 
    if add_err != nil {
        fmt.Println("AddPet Error:", add_err)
    } else {
        fmt.Println("AddPet Response:", add_r)
    }


    add_r, add_err = client.PetApi.AddPet(context.Background(), coffee) 
    if add_err != nil {
        fmt.Println("AddPet Error:", add_err)
    } else {
        fmt.Println("AddPet Response:", add_r)
    }

   // Multiple tags can be provided with comma separated strings. 
   // * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

    //func (a *PetApiService) 
          //FindPetsByTags(ctx context.Context, tags []string) ([]Pet, *http.Response, error) 

    pets_, r, err := client.PetApi.FindPetsByTags(context.Background(),
                         []string {"labradoodle"})

    // Test error
    if err != nil {
        fmt.Println("Pets by Tags Error:", err)
    }

    // Test HTTP Response code
    if r.StatusCode == 200 {
        fmt.Printf("pets=%v\n", pets_)
        fmt.Printf("v=%v\n", r)
    }
}

