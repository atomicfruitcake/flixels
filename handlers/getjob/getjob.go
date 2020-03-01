package getjob

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/atomicfruitcake/flixels/constants"
)

// Handler HTTP Request handler for getting the status encode jobs
func Handler(w http.ResponseWriter, r *http.Request) {
	var jr constants.JobReq
	err := json.NewDecoder(r.Body).Decode(&jr)
	if err != nil {
		msg := fmt.Sprintf("Error decoding request body due to %v", err)
		log.Print(msg)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
