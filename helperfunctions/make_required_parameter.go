package helperfunctions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string
}

func MakeRequiredParameterTypeString(parameter string, parameterName string, w http.ResponseWriter)(isParameterEmpty bool){
	isParameterEmpty = (parameter=="")
	if isParameterEmpty{
		var res Response
		res.Message = fmt.Sprintf("Parameter %s is required",parameterName)
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Println("Error when checking the required Parameters")
		}
	}
	return isParameterEmpty
}