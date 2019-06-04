package main

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/patientplatypus/project/requests"
)

func main() {

	// define mux router
	r := mux.NewRouter()

	// routing
	r.HandleFunc("/pets", requests.PetsGET).Methods("GET")
	r.HandleFunc("/pets", requests.PetsPOST).Methods("POST")
	r.HandleFunc("/pets/{id:[0-9]+}", requests.PetGETone).Methods("GET")

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// serve hot and fresh 15ms or less
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))

}