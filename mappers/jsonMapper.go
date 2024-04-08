package mappers

import (
	"encoding/json"
	"fmt"
)

type provider interface {
	GetByteData() ([]byte, error)
}

type JsonMapper[T any] struct{}

func (jm *JsonMapper[T]) MapValue(p provider, t *T) error {
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
