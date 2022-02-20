package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/darragh-downey/stanley/pkg/app"
)

// JSONHandler handles JSON requests to Stanley and unmarshalls the given request
// data into a StanleyReqPayload struct.
// Then it will return an array of StanleyRes structs of titles with drm content available
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	// necessary to signal to all goroutines that we're done and avoid deadlock
	done := make(chan interface{})
	defer close(done)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprintf(w, "{\"error\": \"Could not decode request: JSON parsing failed\"}", nil)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not decode request: Malformed request body"})
		return
	}

	response, err := app.Parser(done, body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "payload: %v\n", response)
}
