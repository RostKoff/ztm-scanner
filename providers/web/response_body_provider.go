package web

import (
	"fmt"
	"io"
	"net/http"
)

type ResponseBodyProvider struct {
	Url string
}

// Makes GET Request to the specified URL, retrieves the data stored in the response body and returns it as an array of bytes.
func (bp *ResponseBodyProvider) GetByteData() ([]byte, error) {
	resp, err := http.Get(bp.Url)
	if err != nil {
		return nil, fmt.Errorf("failed GetResponseBody: %w", err)
	}
	if resp.StatusCode == 400 {
		return nil, fmt.Errorf("failed GetResponseBody: Bad Request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed GetResponseBody: %w", err)
	}

	return body, err
}
