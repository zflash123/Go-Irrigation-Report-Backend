package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

func GlobalDosPreventer(next http.Handler) http.Handler {
	//This middleware only limit the size of the request Body
	//It is because the request Header size limiter already handled by the net/http by default
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		// Get HTTP_REQUEST_LIMIT_IN_MB and convert it from MB to bytes int64
		var requestSizeLimit string
		var requestSizeLimitInInt int
			//Type assertion to change data type from any to int64
		requestSizeLimit = viper.Get("HTTP_REQUEST_LIMIT_IN_MB").(string)
		requestSizeLimitInInt, _ = strconv.Atoi(requestSizeLimit)
		requestSizeLimitInBytes := int64(requestSizeLimitInInt) << 20
		//HTTP Request Body Size Limiter
		dataRBody, _ := io.ReadAll(r.Body)
		newRBody := io.LimitReader(bytes.NewBuffer(dataRBody), requestSizeLimitInBytes)
			//Detect the error of reading above the requestSizeLimitInBytes
		var err error
		_, err = io.ReadFull(newRBody, dataRBody)
		r.Body = io.NopCloser(bytes.NewBuffer(dataRBody))
		log.Println("Err inside if ct x-www-form= ",err)
			//Handling if error is true
		if err == io.ErrUnexpectedEOF{
			http.Error(w, "Error: DoS Prevented", 413)
			return
		}
		// If there is no error then continue the serve of HTTP with r and w
		next.ServeHTTP(w, r)
	})
}