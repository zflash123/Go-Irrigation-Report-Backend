package helperfunctions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
)

type Response struct {
	Message string
}

func MakeRequiredParameter(parameter any, parameterName string, w http.ResponseWriter)(isParameterEmpty bool){
	assertedParameter := reflect.ValueOf(parameter)
	assertedParameterType := assertedParameter.Type().Name()

	if assertedParameterType == "string"{
		isParameterEmpty = (assertedParameter.String()=="")
	} else{
		isTypeIntOrFloat, _ :=regexp.MatchString("(int.*)|(float.*)", assertedParameterType)
		if isTypeIntOrFloat {
			isParameterEmpty = assertedParameter.IsZero()
		} else {
			isParameterEmpty = !assertedParameter.IsValid()
		}
	}
	if isParameterEmpty{
		var res Response
		res.Message = fmt.Sprintf("Parameter %s is required",parameterName)
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("Error when checking the required Parameters")
		}
		return isParameterEmpty
	} else{
		return false
	}
}