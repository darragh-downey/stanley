package app

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/darragh-downey/stanley/pkg/model"
)

func ConcurrentParser(done chan interface{}, stream []byte) ([]model.StanleyResponse, error) {
	req, errCh := genConcRequest(stream)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-errCh:
				return
			case r := <-req:
				fmt.Println(r)
			}
		}
	}()
	return nil, <-errCh
}

func genConcRequest(stream []byte) (<-chan model.StanleyRequestPayload, <-chan error) {
	req := make(chan model.StanleyRequestPayload)
	errCh := make(chan error)

	jsonData := strings.NewReader(string(stream))
	decoder := json.NewDecoder(jsonData)

	go func() {
		defer close(req)
		defer close(errCh)
		_, err := decoder.Token()
		if err != nil {
			errCh <- fmt.Errorf("Could not decode request: JSON missing opening bracket")
		}

		for decoder.More() {
			// decode an array value
			var payload model.StanleyRequestPayload
			err := decoder.Decode(&payload)
			if err != nil {
				errCh <- fmt.Errorf("Could not decode request: Could not create request struct")

				// potentially metadata
				var meta map[string]interface{}
				err = decoder.Decode(&meta)
				if err != nil {
					errCh <- fmt.Errorf("Could not decode request: Unknown fields %v", err)
				}
			}

			req <- payload
		}

		_, err = decoder.Token()
		if err != nil {
			errCh <- fmt.Errorf("Could not decode request: JSON missing closing bracket")
		}
	}()

	return req, errCh
}

func createPayload(done chan interface{}, stream []byte) (chan model.StanleyRequest, chan error) {
	req := make(chan model.StanleyRequest)
	errCh := make(chan error)

	jsonData := strings.NewReader(string(stream))
	decoder := json.NewDecoder(jsonData)
	go func() {
		defer close(req)
		defer close(errCh)
		for {
			var request model.StanleyRequest
			if err := decoder.Decode(&request); err == io.EOF {
				break
			} else if err != nil {
				// check that we see take, skip and totalRecords else signal error
				var meta map[string]interface{}
				err = decoder.Decode(&meta)
				if err != nil {
					errCh <- fmt.Errorf("Could not decode request: Could not create request struct")
				}
			}
			req <- request
		}
	}()
	return req, errCh
}

// FilterDRM parses a list of requests and returns a list of
// responses containing DRM == true
func filterDRM(done chan interface{}, errChan chan error, stream chan model.StanleyRequest) <-chan model.StanleyResponse {
	res := make(chan model.StanleyResponse)
	go func() {
		defer close(res)
		select {
		case <-done:
			return
		case req := <-stream:
			if req.Drm {
				response, err := model.CreateResponse(req)
				if err != nil {
					errChan <- err
				}
				res <- *response
			}
		case <-errChan:
			return
		}
	}()
	return res
}
