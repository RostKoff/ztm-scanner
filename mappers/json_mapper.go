package mappers

import (
	"encoding/json"
	"fmt"
)

type Provider interface {
	// Retrieves data from a specific source as an array of bytes.
	GetByteData() ([]byte, error)
}

type JsonMapper[T any] struct{}

// Retrieves JSON data from the provider and stores it in the value pointed to by 't'.
func (jm *JsonMapper[T]) MapValue(p Provider, t *T) error {
	data, err := p.GetByteData()
	if err != nil {
		return fmt.Errorf("failed MapValue: %w", err)
	}
	if !json.Valid(data) {
		return fmt.Errorf("failed MapValue: Json is not valid")
	}

	err = json.Unmarshal(data, t)
	return err
}
