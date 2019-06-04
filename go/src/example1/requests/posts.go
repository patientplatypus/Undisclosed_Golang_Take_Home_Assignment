package requests

import(
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
	"encoding/json"
	"github.com/patientplatypus/project/state"
	"github.com/patientplatypus/project/responses"
	"github.com/patientplatypus/project/controller"
)


func PetsPOST(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside PetsPOST")
	decoder := json.NewDecoder(req.Body)
	var data state.Pet
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		responses.ERRORresponse(w, req, err.Error())
	}else{
		controller.AddOne(w, req, data)
	}
}