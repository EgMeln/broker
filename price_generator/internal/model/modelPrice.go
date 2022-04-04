// Package model contain model of struct
package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Price struct that contain record info about Price
type Price struct {
	ID       uuid.UUID
	Ask      float64
	Bid      float64
	Symbol   string
	DoteTime string
}

// MarshalBinary marshal currency to byte
func (pr Price) MarshalBinary() ([]byte, error) {
	return json.Marshal(pr)
}
