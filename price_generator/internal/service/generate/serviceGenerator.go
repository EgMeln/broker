// Package generate to generate price map
package generate

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/EgMeln/broker/price_generator/internal/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const timeFormat = "2006-01-02 15:04:05.00"

// Generator contains generated bid fields
type Generator struct {
	Prices map[string]interface{}
}

// NewGenerator generates bid fields
func NewGenerator() *Generator {
	mapPrices := make(map[string]interface{})
	mapPrices["ALROSA"] = &model.Price{
		ID:       uuid.UUID{},
		Symbol:   "ALROSA",
		Bid:      90.23,
		Ask:      88.37,
		DoteTime: time.Now().Format(timeFormat),
	}
	mapPrices["Aeroflot"] = &model.Price{
		ID:       uuid.UUID{},
		Symbol:   "Aeroflot",
		Bid:      30.96,
		Ask:      30.70,
		DoteTime: time.Now().Format(timeFormat),
	}
	mapPrices["Akron"] = &model.Price{
		ID:       uuid.UUID{},
		Symbol:   "Akron",
		Bid:      14.176,
		Ask:      14.018,
		DoteTime: time.Now().Format(timeFormat),
	}
	return &Generator{
		Prices: mapPrices,
	}
}

// GeneratePrices generates bid prices
func (gen *Generator) GeneratePrices() {
	for key, value := range gen.Prices {
		a, err := rand.Int(rand.Reader, big.NewInt(2))
		if err != nil {
			log.Errorf("rand int error %v", err)
		}
		var coeff float64
		if a.Int64() == 0 {
			coeff = 1.021
		} else {
			coeff = 0.98
		}
		bid := &model.Price{
			Symbol:   key,
			Ask:      value.(*model.Price).Ask * coeff,
			Bid:      value.(*model.Price).Bid * coeff,
			DoteTime: time.Now().Format(timeFormat),
		}
		gen.Prices[key] = bid
		log.Info(gen.Prices[key])
	}
}
