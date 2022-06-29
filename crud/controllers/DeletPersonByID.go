package controllers

import (
	"crud/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//DeletPersonByID delete's person with specific ID
func DeletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	id, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		w.Write([]byte("error in parsingcontacts by id :"))
	}
	result := database.Connector.Where("id = ?", id).Delete(&person)
	if result.Error != nil {
		w.Write([]byte("error in deleting contacts:"))

	} else {

		w.WriteHeader(http.StatusNoContent)
	}

}
