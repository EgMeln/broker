package model

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Price struct {
	ID       uuid.UUID
	Ask      float64
	Bid      float64
	Symbol   string
	DoteTime string
}

func (pr Price) MarshalBinary() ([]byte, error) {
	return json.Marshal(pr)
}
