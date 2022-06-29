package controllers

import (
	"crud/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error in read body:"))
	}

	json.Unmarshal(requestBody, &person)
	result := database.Connector.Save(&person)
	if result.Error != nil {
		w.Write([]byte("error in getting contacts:"))

	} else {
		w.Write([]byte("Contact List : \n"))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(persons)

	}

}
