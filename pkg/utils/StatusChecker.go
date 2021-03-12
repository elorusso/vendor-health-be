package utils

import (
	"errors"
	"net/http"
	"time"
)

type StatusChecker struct {
}

type StatusCheckResponse struct {
	URL        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Duration   int    `json:"duration"`
	Date       int    `json:"date"`
}

func NewStatusChecker() *StatusChecker {
	return &StatusChecker{}
}

func (checker StatusChecker) CheckStatus(url string) (*StatusCheckResponse, error) {

	//check param
	if len(url) == 0 {
		return nil, errors.New("No url provided")
	}

	startTime := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	duration := time.Now().Sub(startTime) //in nanoseconds

	checkResp := &StatusCheckResponse{
		URL:        url,
		StatusCode: resp.StatusCode,
		Duration:   int(duration / 1000000), //convert to miliseconds
		Date:       int(time.Now().Unix()),
	}

	return checkResp, nil
}
