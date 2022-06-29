package controllers

import (
	"crud/database"
	"crud/entity"
	"encoding/json"
	"net/http"
)

var persons []entity.Person

//GetAllPerson get all person data
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := database.Connector.Find(&persons)

	if result.Error != nil {
		w.Write([]byte("error in getting contacts:"))

	} else {
		w.Write([]byte("Contact List : \n"))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(persons)

	}
}
