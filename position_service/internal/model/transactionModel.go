package model

import "github.com/google/uuid"

type Transaction struct {
	ID         uuid.UUID
	PriceOpen  float64
	IsBay      bool
	Symbol     string
	PriceClose float64
}
type GeneratedPrice struct {
	ID       uuid.UUID
	Ask      float64
	Bid      float64
	Symbol   string
	DoteTime string
}
