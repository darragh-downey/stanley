package app

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/darragh-downey/stanley/pkg/model"
)

type Responses struct {
	Responses []model.StanleyResponse
	Error     error
}

type Response struct {
	Response model.StanleyResponse
	Error    error
}
type Request struct {
	Request model.StanleyRequest
	Error   error
}

func ConcurrentParser(done chan interface{}, stream []byte) Responses {
	req := make(chan model.StanleyRequest)
	defer close(req)

	pipeline := filterDRM(done, genConcRequest(done, stream))

	payload := model.CreatePayload()

	var e error
	for res := range pipeline {
		if res.Error != nil {
			e = res.Error
		}
		payload.Add(res.Response)
	}

	return Responses{payload.Responses, e}
}

func genConcRequest(done chan interface{}, stream []byte) <-chan Request {
	req := make(chan Request)

	jsonData := strings.NewReader(string(stream))
	decoder := json.NewDecoder(jsonData)

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("Skipping token")
		}
		// not the token we're looking for

		delim, _ := t.(json.Delim)

		if delim == json.Delim('[') {
			break
		}
	}

	go func() {
		defer close(req)
		for decoder.More() {
			// inside the payload array
			check_decoder := decoder
			if err := checkDuplicateKeys(check_decoder, nil); err != nil {
				// skip this one
				// continue
			}

			var payload model.StanleyRequest
			err := decoder.Decode(&payload)
			request := Request{payload, err}
			if err == io.EOF {
				req <- request
				return
			} else if err != nil {
				request.Error = fmt.Errorf("Could not decode request: Could not create payload struct due to malformed JSON: %v", err)
			}
			req <- request
		}
		fmt.Println("finished")
		done <- struct{}{}
	}()

	return req
}

// FilterDRM parses a list of requests and returns a list of
// responses containing DRM == true
func filterDRM(done chan interface{}, request <-chan Request) <-chan Response {
	res := make(chan Response)
	go func() {
		defer close(res)
		for {
			select {
			case <-done:
				return
			case req := <-request:
				if req.Request.Drm {
					resp, err := model.CreateResponse(req.Request)
					response := Response{*resp, err}
					res <- response
				}
			}
		}
	}()
	return res
}

func checkDuplicateKeys(d *json.Decoder, res map[string]int) error {
	// Get next token from JSON
	t, err := d.Token()
	if err != nil {
		return fmt.Errorf("Could not decode request: Invalid JSON")
	}

	// failing here for closing bracket
	delim, ok := t.(json.Delim)

	// There's nothing to do for simple values (strings, numbers, bool, nil)
	if !ok {
		return nil
	}

	switch delim {
	case '{':
		keys := make(map[string]int)
		for d.More() {
			// Get field key
			t, err := d.Token()
			if err != nil {
				return err
			}
			key := t.(string)

			// There already exists a key in this object
			if keys[key] >= 1 {
				// actually want to return an error here!
				return fmt.Errorf("Could not decode request: Duplicate key in JSON object: %s", key)
			}
			keys[key] += 1

			// Check value
			if err := checkDuplicateKeys(d, res); err != nil {
				return err
			}
		}
		// Consume trailing }
		if _, err := d.Token(); err != nil {
			return err
		}

	case '[':
		i := 0
		for d.More() {
			if err := checkDuplicateKeys(d, res); err != nil {
				return err
			}
			i++
		}
		// Consume trailing ]
		if _, err := d.Token(); err != nil {
			return err
		}

	}
	return nil
}

func testLevel(decoder *json.Decoder, keys map[string]int, level int) error {
	// gather all keys on current level and add to map
	if decoder.More() {
		t, err := decoder.Token()
		if err == io.EOF {
			// end of input stream
			return nil
		} else if err != nil {
			// Token returns nil or EOF at end of input so this is a legitimate error
			return err
		}

		// test if we have a delimiter character
		_, ok := t.(json.Delim)

		// There's nothing to do for simple values (strings, numbers, bool, nil)
		if !ok {
			key := t.(string)                                // human readable key
			m_key := fmt.Sprintf("%s_%d", t.(string), level) // unique key per level
	
			if keys[m_key] >= 1 {
				return fmt.Errorf("Could not decode request: Duplicate key in JSON object: %s", key)
			}
			keys[m_key] += 1
		}


		return testLevel(decoder, keys, level+1)
	}
	// nothing more to take so we've reached the end of the JSON object
	return nil
}
