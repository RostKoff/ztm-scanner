package client

import (
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       []byte
}

// Returns the body and status code of the response after making GET request to the passed URL.
func GetResponse(url string) (Response, error) {
	respHolder := Response{}
	resp, err := http.Get(url)
	if err != nil {
		return respHolder, fmt.Errorf("failed GetResponseBody: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed GetResponseBody: %w", err)
	} else {
		respHolder.StatusCode = resp.StatusCode
		respHolder.Body = body
	}

	return respHolder, err
}
