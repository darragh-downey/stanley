package app

import (
	"encoding/json"
	"fmt"

	"github.com/darragh-downey/stanley/pkg/model"
)

func Parser(done chan interface{}, stream []byte) ([]model.StanleyResponse, error) {
	reqs := make(chan model.StanleyRequest)
	errChan := make(chan error)
	s := make(chan []byte)

	req := createRequest(done, errChan, s)
	resp := filterDRM(done, errChan, reqs)

	responses := make([]model.StanleyResponse, 0, 10)

	for r := range resp {
		responses = append(responses, r)
	}

	return responses, <-errChan
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
		}
	}()
	return res
}

func createRequest(done chan interface{}, errChan chan error, stream chan []byte) <-chan model.StanleyRequest {
	res := make(chan model.StanleyRequest)
	go func() {
		defer close(res)
		select {
		case <-done:
			return
		case s := <-stream:
			payload := &model.StanleyRequestPayload{}
			err := json.Unmarshal([]byte(s), payload)
			if err != nil {
				errChan <- fmt.Errorf("Could not decode request: JSON parsing failed")
			}
			for _, r := range payload.Requests {
				res <- r
			}
		case <-errChan:
			return
		}
	}()
	return res
}
