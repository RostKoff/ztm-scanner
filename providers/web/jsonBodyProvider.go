package web

import (
	"fmt"
	"io"
	"net/http"
)

type ResponseBodyProvider struct {
	Url string
}

func (bp *ResponseBodyProvider) GetByteData() ([]byte, error) {
	resp, err := http.Get(bp.Url)
	if err != nil {
		return nil, fmt.Errorf("failed GetResponseBody: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed GetResponseBody: %w", err)
	}

	return body, err
}
