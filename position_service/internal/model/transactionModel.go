// Package model contain model of struct
package model

import (
	"github.com/google/uuid"
)

// Transaction struct that contain record info about transaction
type Transaction struct {
	ID         uuid.UUID `json:"id_"`
	PriceOpen  float64   `json:"price_open"`
	IsBay      bool      `json:"is_bay"`
	Symbol     string    `json:"symbol"`
	PriceClose float64   `json:"price_close"`
	BayBy      string    `json:"bay_by"`
}

// GeneratedPrice struct that contain record info about new price
type GeneratedPrice struct {
	ID       uuid.UUID
	Ask      float64
	Bid      float64
	Symbol   string
	DoteTime string
}
