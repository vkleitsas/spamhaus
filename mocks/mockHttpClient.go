package mocks

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HTTPClientMock struct {
}

// By mocking the Do function we can have preset responses to our test data
// This way, we can validate the logic of the worker without external dependencies
func (c *HTTPClientMock) Do(req *http.Request) (*http.Response, error) {
	if req.URL.String() == "http://example.com/1" {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}, nil
	} else if req.URL.String() == "http://example.com/2" {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}, nil
	} else if req.URL.String() == "http://example.com/3" {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}, nil
	}
	return nil, fmt.Errorf("invalid url")
}

func NewMockHTTPClient() *HTTPClientMock {
	return &HTTPClientMock{}
}
