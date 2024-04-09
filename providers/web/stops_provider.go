package web

import (
	"bufio"
	"fmt"
	"net/http"
)

const (
	stopsUrl = "https://ckan.multimediagdansk.pl/dataset/c24aa637-3619-4dc2-a171-a23eec8f2172/resource/4c4025f0-01bf-41f7-a39f-d156d201b82b/download/stops.json"
)

type StopsProvider struct{}

// Gets the latest data on public transport stops provided by ZTM and returns it as an array of bytes.
func (s *StopsProvider) GetByteData() ([]byte, error) {
	resp, err := http.Get(stopsUrl)
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
