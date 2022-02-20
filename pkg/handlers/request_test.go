package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/darragh-downey/stanley/pkg/handlers"
	"github.com/darragh-downey/stanley/pkg/model"
)

// Some quick reading on testing handlers
// https://github.com/gorilla/mux#testing-handlers

func TestSimpleRequest(t *testing.T) {
	tt := []struct {
		reqs     model.StanleyRequest
		expected model.StanleyResponse
	}{
		{
			model.StanleyRequest{},
			model.StanleyResponse{},
		},
		{
			model.StanleyRequest{},
			model.StanleyResponse{},
		},
		{
			model.StanleyRequest{},
			model.StanleyResponse{},
		},
	}

	for _, testCase := range tt {

		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.JSONHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Failed with status: %v %v\n", status, http.StatusOK)
		}

		res := &model.StanleyResponse{}

		err = json.Unmarshal(rr.Body.Bytes(), res)
		if err != nil {
			t.Error(err)
		}

		if *res != testCase.expected {
			t.Errorf("Failed parsing the request: %v %v\n", rr.Body.String(), testCase.expected)
		}
	}
}
