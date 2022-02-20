package model

type StanleyResponsePayload struct {
	Responses []StanleyResponse `json:"response"`
}

type StanleyResponse struct {
	Image string `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

func CreateResponse(request StanleyRequest) (*StanleyResponse, error) {
	image := request.Image.ShowImage
	slug := request.Slug
	title := request.Title

	// validate that the given request is valid
	// otherwise build an error message detailing the failed validation

	return &StanleyResponse{
		Image: image,
		Slug:  slug,
		Title: title,
	}, nil
}
