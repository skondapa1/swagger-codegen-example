package server

import (
   "fmt"
   "reflect"
   pets "swagger-codegen-example/pets/go"
)

var  petsArr  []pets.Pet

func Find (slice interface{}, f func(value interface{}) bool) int {
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
func InsertPetToArr (pet pets.Pet) error {
    petsArr= append(petsArr, pet)
    return nil
}

func FindPetByIdFromArr (petId int64)  *pets.Pet {
    idx :=  Find(petsArr, func(value interface{}) bool {
       return value.(pets.Pet).Id == petId
    })

    if (idx < 0) {
        return nil
    } else {
        fmt.Println("GetMyPet", petsArr[idx])
        return &petsArr[idx]
    }
}

/**
 * DeletePetById_
 *
 * Delete a pet by Id
 */
func DeletePetByIdFromArr (petId int64) int {

    idx :=  Find(petsArr, func(value interface{}) bool {
       return value.(pets.Pet).Id ==  petId
    })

    if (idx >= 0) {
        // changes order of pets
        petsArr[idx] = petsArr[len(petsArr)-1]
        petsArr = petsArr[:len(petsArr)-1]
    }

    return idx
}

