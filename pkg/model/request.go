package model

type StanleyRequestPayload struct {
	Requests     []StanleyRequest
	Skip         int `json:"skip"`
	Take         int `json:"take"`
	TotalRecords int `json:"totalRecords"`
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
	Seasons       []string       `json:"seasons"`
	Slug          string         `json:"slug"`
	Title         string         `json:"title"`
	TvChannel     string         `json:"tvChannel"`
}

func CreateRequest(s []byte) (*StanleyRequest, error) {

	return &StanleyRequest{}, nil
}
