package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/darragh-downey/stanley/pkg/util"
)

type StanleyRequestPayload struct {
	Requests     []StanleyRequest `json:"payload"`
	Skip         int              `json:"skip"`
	Take         int              `json:"take"`
	TotalRecords int              `json:"totalRecords"`
}

type StanleyRequest struct {
	Country       string         `json:"country"`
	Description   string         `json:"description"`
	Drm           bool           `json:"drm"`
	EpisodeCount  int            `json:"episodeCount"`
	Genre         string         `json:"genre"`
	Image         StanleyImage   `json:"image"`
	Language      string         `json:"language"`
	NextEpisode   StanleyEpisode `json:"nextEpisode"`
	PrimaryColour string         `json:"primaryColour"`
	Seasons       []Season       `json:"seasons"`
	Slug          string         `json:"slug"`
	Title         string         `json:"title"`
	TvChannel     string         `json:"tvChannel"`
}

func CreateRequest(s []byte) (*StanleyRequest, error) {

	return &StanleyRequest{}, nil
}

func (s *StanleyRequest) UnmarshalJSON(data []byte) error {
	keys, err := util.DetectDuplicateKeys(data)
	if err != nil {
		return err
	}

	for k, v := range keys {
		if v > 1 {
			key := strings.Split(k, "_")
			return fmt.Errorf("Unable to unmarshal JSON Request object - Key %s at level %s appears %d times", key[0], key[1], v)
		}
	}

	// https://stackoverflow.com/questions/43176625/call-json-unmarshal-inside-unmarshaljson-function-without-causing-stack-overflow/43178272#43178272
	type request2 StanleyRequest
	if err := json.Unmarshal(data, (*request2)(s)); err != nil {
		return err
	}

	return nil
}
