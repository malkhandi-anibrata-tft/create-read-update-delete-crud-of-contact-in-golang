package controllers

import (
	"crud/database"
	"crud/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var person entity.Person

//CreatePerson creates person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error in read body:"))
	}

	json.Unmarshal(requestBody, &person)

	result := database.Connector.Create(person)

	if result.Error != nil {
		w.Write([]byte("error in creating contacts:"))

	} else {

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(persons)

	}

}
