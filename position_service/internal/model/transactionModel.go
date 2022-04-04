// Package model contain model of struct
package model

import "github.com/google/uuid"

// Transaction struct that contain record info about transaction
type Transaction struct {
	ID         uuid.UUID
	PriceOpen  float64
	IsBay      bool
	Symbol     string
	PriceClose float64
}

// GeneratedPrice struct that contain record info about new price
type GeneratedPrice struct {
	ID       uuid.UUID
	Ask      float64
	Bid      float64
	Symbol   string
	DoteTime string
}
