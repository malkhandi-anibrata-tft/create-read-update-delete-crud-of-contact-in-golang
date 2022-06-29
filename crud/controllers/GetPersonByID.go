package controllers

import (
	"crud/database"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// var person entity.Person

//GetPersonByID returns person with specific ID
func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	result := database.Connector.First(&person, key)
	if result.Error != nil {
		w.Write([]byte("error in getting contacts by id :"))

	} else {
		w.Write([]byte("Contact : \n"))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(persons)

	}

}
