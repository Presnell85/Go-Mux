package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// a struct is similar to an array in OOP but the main difference that a struct can hold varibles
// of different data types, lets think of it as a container

type Medication struct {
	ID      int    `json:"id"`
	Brand   string `json:"brand"`
	Generic string `json: "generic"`
	Dosage  int    `json: "dosage"`
}

var medications []Medication

func main() {
	r := mux.NewRouter()

	medications = append(medications, Medication{ID: 1, Brand: "phizer", Generic: "", Dosage: 3},
		Medication{ID: 2, Brand: "watson", Generic: "", Dosage: 2},
		Medication{ID: 3, Brand: "Aztra-Zentica", Generic: "", Dosage: 1})

	r.HandleFunc("/medications", getMedications).Methods("GET")
	r.HandleFunc("/medications/{id}", getMedication).Methods("GET")
	r.HandleFunc("/medications", addMedication).Methods("POST")
	r.HandleFunc("/medications", updateMedications).Methods("PUT")
	r.HandleFunc("/medications", removeMedications).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}

//this is how you return all medications
func getMedications(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(medications)
}

//this is how you select a single endpoint by id
func getMedication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"]) //this is just converting the id to an int from a string
	for _, medication := range medications {
		if medication.ID == i {
			json.NewEncoder(w).Encode(&medication)
		}
	}
}

// this is how you add a medication
func addMedication(w http.ResponseWriter, r *http.Request) {
	var medication Medication
	json.NewDecoder(r.Body).Decode(&medication)   //this is where you are actually creating the new json content
	medications = append(medications, medication) // this is the format that you are creating
	json.NewEncoder(w).Encode(medications)        // this is showing you the new medication
}

//update a medication by searching all medications id and looking for a change
func updateMedications(w http.ResponseWriter, r *http.Request) {
	var medication Medication
	json.NewDecoder(r.Body).Decode(&medication)
	for i, item := range medications {
		if item.ID == medication.ID {
			medications[i] = medication
		}
	}

}

//remove a medication
func removeMedications(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"]) //convert to int
	for i, item := range medications {
		if item.ID == id {
			medications = append(medications[:i], medications[i+1:]...)
		}
	}
}
