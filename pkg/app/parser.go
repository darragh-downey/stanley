package app

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/darragh-downey/stanley/pkg/model"
)

func LinearParser(stream []byte) (model.StanleyResponsePayload, error) {
	if len(stream) == 0 {
		return model.StanleyResponsePayload{}, fmt.Errorf("Could not decode request: Empty request")
	}

	requests, err := genRequest(stream)
	if err != nil {
		return model.StanleyResponsePayload{}, err
	}

	responses, err := genResponses(requests)
	if err != nil {
		return model.StanleyResponsePayload{}, err
	}

	return responses, nil
}

func genRequest(stream []byte) ([]model.StanleyRequest, error) {
	jsonData := strings.NewReader(string(stream))
	decoder := json.NewDecoder(jsonData)

	requests := make([]model.StanleyRequest, 0, 10)

	for {
		var payload model.StanleyRequestPayload
		err := decoder.Decode(&payload)
		if err == io.EOF {
			break
		} else if err != nil {
			return requests, fmt.Errorf("Could not decode request: Could not create payload struct due to malformed JSON: %v", err)
		}
		requests = payload.Requests
		// fmt.Printf("%+v\n", requests)
	}

	return requests, nil
}

func genResponses(requests []model.StanleyRequest) (model.StanleyResponsePayload, error) {
	payload := model.CreatePayload()
	for _, request := range requests {
		if request.Drm && request.EpisodeCount > 0 {
			response, err := model.CreateResponse(request)
			if err != nil {
				return model.StanleyResponsePayload{}, err
			}
			payload.Add(*response)
		}
	}
	return *payload, nil
}
