package test

import (
	"bytes"
	"crud/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetAllPerson(t *testing.T) {
	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetAllPerson)
	handler.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusOK, rr.Code)

}

func TestGetPersonByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/get/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetPersonByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check  whether the response body is as expected.
	expected := `{"ID":1,"FirstName":"john","LastName":"wick",Phone":9779993450}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreatePerson(t *testing.T) {

	var jsonStr = []byte(`{"ID":4,"FirstName":"xyz","LastName":"pqr","phone":1234567890}`)

	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreatePerson)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ID":4,"FirstName":"xyz","LastName":"pqr","Phone":"1234567890"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func UpdatePersonByID(t *testing.T) {

	var jsonStr = []byte(`{"ID":4,"FirstName":"xyz change","LastName":"pqr","Phone":"1234567890"}`)

	req, err := http.NewRequest("PUT", "/update/{id}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.UpdatePersonByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ID":4,"FirstName":"xyz change","LastName":"pqr","Phone":"1234567890"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func DeletPersonByID(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delete/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.DeletPersonByID)
	handler.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusOK, rr.Code)

}
