package utils

import (
	"assignment/domain"
	"assignment/mocks"
	"testing"
)

func TestDownloadUrls_OK(t *testing.T) {
	client := mocks.NewMockHTTPClient()

	urlDownload := NewUrlDownload(client)

	urls := []domain.URLEntry{
		{WebsiteURL: "http://example.com/1"},
		{WebsiteURL: "http://example.com/2"},
		{WebsiteURL: "http://example.com/3"},
	}
	_, successCount, failureCount := urlDownload.DownloadUrls(urls)

	if successCount != len(urls) {
		t.Errorf("Expected successful downloads count to be %d, but got %d", len(urls), successCount)
	}

	// Assert failure count is 0
	if failureCount != 0 {
		t.Errorf("Expected failure count to be 0, but got %d", successCount)
	}
}

func TestDownloadUrls_OneDownloadFailing(t *testing.T) {
	client := mocks.NewMockHTTPClient()

	urlDownload := NewUrlDownload(client)

	urls := []domain.URLEntry{
		{WebsiteURL: "http://example.com/1"},
		{WebsiteURL: "http://example.com/2"},
		{WebsiteURL: "http://badexample.com/3"},
	}
	_, successCount, failureCount := urlDownload.DownloadUrls(urls)

	if successCount != len(urls)-1 {
		t.Errorf("Expected successful downloads count to be %d, but got %d", len(urls), successCount)
	}

	// Assert failure count is 1
	if failureCount != 1 {
		t.Errorf("Expected failure count to be 0, but got %d", successCount)
	}
}
