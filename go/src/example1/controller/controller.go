package controller

import(
	"fmt"
	"net/http"
	"strconv"
	"github.com/patientplatypus/project/state"
	"github.com/patientplatypus/project/responses"
)

func convert64(id string) (int64, string) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		return -1, err.Error()
	}
	return i, ""
}

func FindOne(w http.ResponseWriter, req *http.Request, id string){
	id64, err64 := convert64(id)
	if(len(err64)>0){
		responses.ERRORresponse(w, req, err64)
	}else{
		matchfound := false
		for i := range state.Petarray {
			if state.Petarray[i].ID == id64 {
				fmt.Println("found match!")
				matchfound = true
				fmt.Println(state.Petarray[i])
				responses.GETresponse(w, req, "one", i)
			}else if matchfound == false{
				fmt.Println("no match found!")
				matchfound = false
			}
		}
		if matchfound == false{
			responses.ERRORresponse(w, req, "no match found!")
		}
	}
}

func checkUniqueID(id int64)bool{
	matchfound := false
	for i := range state.Petarray {
		if state.Petarray[i].ID == id {
			matchfound = true
			return false
		}
	}
	if matchfound == false{
		return true
	}
	return false
}

func AddOne(w http.ResponseWriter, req *http.Request, data state.Pet){
	fmt.Println("inside AddOne")
	uniqueIDCheck := checkUniqueID(data.ID)
	if uniqueIDCheck == true{
		state.Petarray = append(state.Petarray, data)
	}else{

	}
	responses.POSTresponse(w, req)
}