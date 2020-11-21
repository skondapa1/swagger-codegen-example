package swagger

import (
    "net/http"
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "strconv"
    "strings"
)

func AddMyPet(w http.ResponseWriter, r *http.Request) {
    var pet Pet
    // _ = json.NewDecoder(r.Body).Decode(pet)
    // fmt.Println("AddMyPet", r.Body)
    // fmt.Println("AddMyPet", pet)

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
       log.Printf("Error reading body: %v", err)
       http.Error(w, "can't read body", http.StatusBadRequest)
       return
    }
    json.Unmarshal([]byte(body), &pet)
    fmt.Println("_AddMyPet", pet)

    err = InsertPet(pet)
    if err != nil {
       http.Error(w, "Db Insert failed !!", http.StatusBadRequest)
       return
    }

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func DeleteMyPet(w http.ResponseWriter, r *http.Request) {
    var pet Pet
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
       log.Printf("Error reading body: %v", err)
       http.Error(w, "can't read body", http.StatusBadRequest)
       return
    }
    json.Unmarshal([]byte(body), &pet)

    idx := DeletePetById(pet.Id)

    if (idx < 0) {
	    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(http.StatusNotFound)
    } else {
	    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(http.StatusOK)
    }
}

func GetMyPetById(w http.ResponseWriter, r *http.Request) {

    log.Printf("GetMyPetById  path %s", r.URL.Path)
    id := strings.TrimPrefix(r.URL.Path, "/v2/pet/")

    if (id == "")  {
	    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(http.StatusNotFound)
        return
    }

    petId, _ :=  strconv.ParseInt(id,  10, 64)
    log.Printf("GetMyPetById  %v", petId)

    pet := FindPetById(petId)

    if (pet == nil) {
	    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(http.StatusNotFound)
    } else {
        json.NewEncoder(w).Encode(pet)
	    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(http.StatusOK)
    }
}

func FindMyPetsByTags(w http.ResponseWriter, r *http.Request) {

    log.Printf("GetMyPetByTags  path %s", r.URL.Path)

    query := r.URL.Query()
    filters, present := query["tags"] 

    if !present || len(filters) == 0 {
        fmt.Println("tags not present")
	    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(http.StatusNotFound)
        return
    }

    log.Printf("GetMyPetByTag  %v", filters)

    pets, _ := FindPetsByTag(filters)
    if (pets == nil) {
	    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(http.StatusNotFound)
    } else {
        fmt.Printf("-Found pets %v\n", pets)
	    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(pets)
    }
}

