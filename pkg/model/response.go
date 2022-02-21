package model

type StanleyResponsePayload struct {
	Responses []StanleyResponse `json:"response"`
	setOf     map[string]StanleyResponse
}

type StanleyResponse struct {
	Image string `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

func CreatePayload() *StanleyResponsePayload {
	responses := make([]StanleyResponse, 0, 10)
	return &StanleyResponsePayload{
		Responses: responses,
		setOf:     make(map[string]StanleyResponse),
	}
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

func (p *StanleyResponsePayload) Add(resp StanleyResponse) {
	if _, ok := p.setOf[resp.Title]; !ok {
		// doesn't exist in set
		p.setOf[resp.Title] = resp
		p.Responses = append(p.Responses, resp)
	}
}
