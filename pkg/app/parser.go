package app

import (
	"encoding/json"
	"fmt"

	"github.com/darragh-downey/stanley/pkg/model"
)

func Parser(done chan interface{}, stream []byte) ([]model.StanleyResponse, error) {
	payload := &model.StanleyRequestPayload{}

	err := json.Unmarshal(stream, payload)
	if err != nil {
		return nil, fmt.Errorf("Could not decode request: JSON parsing failed")
	}

	reqs := make(chan model.StanleyRequest)

	for _, req := range payload.Requests {
		reqs <- req
	}

	errChan := make(chan error)

	filterDRM(done, reqs, errChan)

	return nil, <-errChan
}

// FilterDRM parses a list of requests and returns a list of
// responses containing DRM == true
func filterDRM(done chan interface{}, stream chan model.StanleyRequest, errChan chan error) <-chan model.StanleyResponse {
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
