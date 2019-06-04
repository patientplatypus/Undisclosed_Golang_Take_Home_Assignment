package responses

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/patientplatypus/project/state"
)

type AllResponse struct {
	Status string
	Petarray []state.Pet
}
type OneResponse struct {
	Status string
	Petone state.Pet
}
type ErrorResponse struct{
	Status string
	Message string
}

func POSTresponse(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside POSTresponse")
	response := AllResponse{Petarray: state.Petarray, Status: "200"}
	json.NewEncoder(w).Encode(response) 
}

func GETresponse(w http.ResponseWriter, req *http.Request, getType string, getIndex int){
	fmt.Println("inside GETresponse")
	if getType == "all"{
		fmt.Println("inside GETresponse ALL")
		response := AllResponse{Petarray: state.Petarray, Status: "200"}
		json.NewEncoder(w).Encode(response) 
	}else{
		fmt.Println("inside GETresponse ONE")
		response := OneResponse{Petone: state.Petarray[getIndex], Status: "200"}
		json.NewEncoder(w).Encode(response) 
	}
}

func ERRORresponse(w http.ResponseWriter, req *http.Request, message string){
	fmt.Println("inside ERRORresponse")
	response := ErrorResponse{Message: message, Status: "400"}
	json.NewEncoder(w).Encode(response) 
}