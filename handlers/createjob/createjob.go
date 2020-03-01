package createjob

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/atomicfruitcake/flixels/mediaconvertclient"
	"github.com/atomicfruitcake/flixels/constants"
)

// Handler HTTP Request handler for create new jobs for requests
func Handler(w http.ResponseWriter, r *http.Request) {
	var j constants.Job
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		msg := fmt.Sprintf("Error decoding request body due to %v", err)
		log.Print(msg)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	j.Status = constants.Processing
	params, err := json.Marshal(j)
	if err != nil {
		msg := fmt.Sprintf("Error marshalling request body due to %v", err)
		log.Print(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	err = mediaconvertclient.CreateEncodeJob(j.S3URL, j.Rendition)
	if err != nil {
		msg := fmt.Sprintf("Error sending job to service due to %v", err)
		log.Print(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
