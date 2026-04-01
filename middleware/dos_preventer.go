package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

func GlobalDosPreventer(next http.Handler) http.Handler {
	//This middleware only limit the size of the request Body
	//It is because the request Header size limiter already handled by the net/http by default
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		var requestSizeLimit string
		var requestSizeLimitInInt int
		//Type assertion to change data type from any to int64
		requestSizeLimit = viper.Get("HTTP_REQUEST_LIMIT_IN_MB").(string)
		requestSizeLimitInInt, _ = strconv.Atoi(requestSizeLimit)
		maxBytesSize := int64(requestSizeLimitInInt) << 20
		//Add HTTP Request Body Size Limiter Handler
		r.Body = http.MaxBytesReader(w, r.Body, maxBytesSize)
		err := r.ParseForm()
		if err != nil{
			var res Response
			res.Message = "Your request is limited for DoS Prevention"
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				log.Println("Error when Encode the res to JSON using w Encoder")
			}
		}
		//Return the new Handler
		next.ServeHTTP(w, r)
	})
}