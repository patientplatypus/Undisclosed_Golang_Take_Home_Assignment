package requests

import(
	"fmt"
	"net/http"
	// "encoding/json"
	"github.com/gorilla/mux"
	"github.com/patientplatypus/project/responses"
	"github.com/patientplatypus/project/controller"
)



func PetsGET(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside PetsGET")
	responses.GETresponse(w, req, "all", -1)
}

func PetGETone(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside PetGETone")
	vars := mux.Vars(req)
	// fmt.Println("value of vars")
	// fmt.Println(vars)
	controller.FindOne(w, req, vars["id"])
}
