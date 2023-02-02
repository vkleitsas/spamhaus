package utils

import (
	"assignment/domain"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type UrlDownload struct {
	client HttpClientInterface
}

func NewUrlDownload(c HttpClientInterface) *UrlDownload {
	return &UrlDownload{
		client: c,
	}
}

type UrlDownloadInterface interface {
	DownloadSingleUrl(r domain.URLEntry) (*domain.URLEntry, error)
	DownloadUrls(urls []domain.URLEntry) ([]time.Duration, int, int)
}

func (u *UrlDownload) DownloadSingleUrl(r domain.URLEntry) (*domain.URLEntry, error) {
	request, err := http.NewRequest(http.MethodGet, r.WebsiteURL, nil)
	if err != nil {
		log.Println("Failed to create request:", err)

	}
	resp, err := u.client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Request failed with status " + fmt.Sprint(resp.StatusCode))
	}
	r.Data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (u *UrlDownload) DownloadUrls(urls []domain.URLEntry) ([]time.Duration, int, int) {
	// Set the concurrency factor to 3
	waitChan := make(chan struct{}, 3)
	results := make(chan time.Duration, len(urls))

	var successCount int
	var failureCount int

	for _, url := range urls {
		// Get a slot of the waitChan
		waitChan <- struct{}{} // Execution is blocked here until a slot of the waitChan is free
		go func(url domain.URLEntry) {
			start := time.Now()
			request, err := http.NewRequest(http.MethodGet, url.WebsiteURL, nil)
			if err != nil {
				log.Println("Failed to create request:", err)
				failureCount++
			}
			// Make a GET request to the URL
			resp, err := u.client.Do(request)
			if err != nil {
				log.Println("Failed to download URL:", err)
				failureCount++
			} else {
				resp.Body.Close()
				successCount++
			}
			// Release a slot of the waitChan
			<-waitChan
			results <- time.Since(start)
		}(url)
	}

	downloadtimes := []time.Duration{}
	// Wait for all goroutines to finish
	for i := 0; i < len(urls); i++ {
		downloadtimes = append(downloadtimes, <-results)
	}
	// Log the results
	log.Println("Download times:", downloadtimes)
	log.Println("Successful downloads:", successCount)
	log.Println("Failed downloads:", failureCount)

	return downloadtimes, successCount, failureCount
}
