package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type reportSegment struct {
	Segment_id string `json:"segment_id"`
	Level      string `json:"level"`
	Note       string `json:"note"`
	Image      string `json:"image"`
}

func CreateReportContentTypeHandler(w http.ResponseWriter, r *http.Request) (reportSegments []reportSegment) {
	var res Response
	//If the req body use JSON data, then decode the JSON to an object
	if r.Header.Get("Content-Type") == "application/json" {
		var reportSegments []reportSegment
		json.NewDecoder(r.Body).Decode(&reportSegments)
		return reportSegments
		//Else If the req body use x-www-form-urlencoded data, so it will parsed as form data
	} else if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		totalKeysPostForm := len(r.PostForm)
		if err != nil {
			log.Fatalln("Something went wrong when ParseForm:\n", err)
		}
		log.Println("len(r.PostForm): ", len(r.PostForm))

		//Check All Required Value of Each Key
		if totalKeysPostForm < 4 {
			w.WriteHeader(http.StatusBadRequest)
			res.Message = "All parameters is required"
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				log.Fatalln(err)
			}
			return nil
		}
		var lengthValuesHighest = 0
		var lengthValuesLowest = 4
		for key, values := range r.PostForm {
			var currentValuesLength = len(values)
			// Check length of each values >3 or not
			if currentValuesLength > 3 {
				w.WriteHeader(http.StatusBadRequest)
				res.Message = fmt.Sprintf("too many values for parameter '%s'", key)
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					log.Fatalln(err)
				}
				return nil
			}

			//In Loop will Return the Highest Length of Array Values [Values must not exceeds 3]
			if currentValuesLength > lengthValuesHighest {
				lengthValuesHighest = currentValuesLength
			}
			if currentValuesLength < lengthValuesLowest {
				lengthValuesLowest = currentValuesLength
			}
		}

		// Check wether the each array length is same or not
		if lengthValuesHighest != lengthValuesLowest {
			w.WriteHeader(http.StatusBadRequest)
			res.Message = "All parameters is required"
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				log.Fatalln(err)
			}
			return nil
		}
		//Dapatkan length
		var lengthPostForm = lengthValuesLowest
		var reportSegments = make([]reportSegment, lengthPostForm)
		for i := 0; i <= lengthPostForm-1; i++ {
			reportSegments[i].Segment_id = r.PostForm["segment_id"][i]
			reportSegments[i].Level = r.PostForm["level"][i]
			reportSegments[i].Note = r.PostForm["note"][i]
			reportSegments[i].Image = r.PostForm["image"][i]
		}
		return reportSegments
	} else {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = "Content-Type not supported"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Fatalln(err)
		}
		return nil
	}
}