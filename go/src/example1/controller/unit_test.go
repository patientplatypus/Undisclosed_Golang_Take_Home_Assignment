package controller

import(
	"fmt"
	"testing"
	// "reflect"
)

//here I am only going to unit test the convert64 function
//notice that this test failing prevents the docker container from compiling. Good, this prevents pushing broken code!

//if we wanted to write integration tests, I would probably make the run.sh file run headless, then write a pure golang function that hit http routes and returned results depending on expected conditions WITHOUT using the "testing" package. This would just overall make it easier I think

func TestConvert64(t *testing.T) {
	inputs := "823982"
	returnVal, returnError := convert64(inputs)
  
	fmt.Println(returnVal)
	fmt.Println(returnError)
	if returnVal < 0 || len(returnError) != 0 {
	  t.Errorf("found an error in unit test!")
	}else{
	  fmt.Println("good return passes")
	}
}  