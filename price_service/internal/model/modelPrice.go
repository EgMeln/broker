// Package model contain model of struct
package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

// GeneratedPrice struct that contain record info about new Price
type GeneratedPrice struct {
	ID       uuid.UUID
	Ask      float64
	Bid      float64
	Symbol   string
	DoteTime string
}

// MarshalBinary marshal currency to byte
func (gen *GeneratedPrice) MarshalBinary() ([]byte, error) {
	return json.Marshal(gen)
}

// UnmarshalBinary unmarshal currency from byte
func (gen *GeneratedPrice) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, gen)
}
