package utils

import "net/http"

// This interface is created to allow us to create a mock of the actual http.Client,
// which is used in the urlDownload tests
type HttpClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}
