package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/darragh-downey/stanley/pkg/app"
)

// JSONLinearHandler handles JSON requests to Stanley and unmarshalls the given request
// data into a StanleyReqPayload struct.
// Then it will return an array of StanleyRes structs of titles with drm content available
func JSONLinearHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprintf(w, "{\"error\": \"Could not decode request: JSON parsing failed\"}", nil)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not decode request: Malformed request body"})
		return
	}

	response, err := app.LinearParser(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	// fmt.Fprintf(w, "%v\n", response)
}

func JSONConcHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan interface{})
	defer close(done)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not decode request: Malformed request body"})
		return
	}

	response, err := app.ConcurrentParser(done, body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "payload: %v\n", response)
}
