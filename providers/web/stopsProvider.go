package web

import (
	"bufio"
	"fmt"
	"net/http"
)

type StopsProvider struct {
	Url string
}

func (s *StopsProvider) GetByteData() ([]byte, error) {
	resp, err := http.Get(s.Url)
	if err != nil {
		return nil, fmt.Errorf("failed GetByteData: %w", err)
	}
	defer resp.Body.Close()

	r := bufio.NewReader(resp.Body)

	// Discarding unnecessary part of the JSON.
	_, err = r.Discard(58)
	if err != nil {
		return nil, fmt.Errorf("failed GetByteData: %w", err)
	}

	// Getting only the most up-to-date list of stops.
	bytesData, err := r.ReadBytes(']')
	if err != nil {
		return nil, fmt.Errorf("failed GetBytesData: %w", err)
	}
	return bytesData, nil
}
