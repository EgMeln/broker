// Package service contains business logic
package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sync"

	"github.com/EgMeln/broker/position_service/internal/model"
	"github.com/EgMeln/broker/position_service/internal/repository"
	"github.com/google/uuid"
)

// PositionService struct for
type PositionService struct {
	rep          repository.PriceTransaction
	generatedMap map[string]*model.GeneratedPrice
	mu           *sync.RWMutex
	positionMap  map[string]map[string]*chan *model.GeneratedPrice
}

// NewPositionService used for setting position services
func NewPositionService(rep *repository.PostgresPrice, priceMap map[string]*model.GeneratedPrice, mute *sync.RWMutex, pos map[string]map[string]*chan *model.GeneratedPrice) *PositionService {
	return &PositionService{rep: rep, generatedMap: priceMap, mu: mute, positionMap: pos}
}

// OpenPosition add record about position
func (src *PositionService) OpenPosition(ctx context.Context, trans *model.Transaction, str string) (*uuid.UUID, error) {
	ch := make(chan *model.GeneratedPrice)
	src.mu.Lock()
	src.positionMap[trans.Symbol][trans.ID.String()] = &ch
	src.mu.Unlock()
	if str == "Ask" {
		go src.getProfitByAsk(ch, trans)
	} else if str == "Bid" {
		go src.getProfitByBid(ch, trans)
	}
	return src.rep.OpenPosition(ctx, trans)
}

// ClosePosition update record about position
func (src *PositionService) ClosePosition(ctx context.Context, closePrice *float64, id *uuid.UUID, str string) (string, error) {
	src.mu.Lock()
	close(*src.positionMap[str][(*id).String()])
	delete(src.positionMap[str], id.String())
	src.mu.Unlock()
	return src.rep.ClosePosition(ctx, closePrice, id)
}

func (src *PositionService) getProfitByAsk(ch chan *model.GeneratedPrice, trans *model.Transaction) {
	for {
		price, ok := <-ch
		if ok {
			log.Printf("For position %v profit if close: %v", trans.ID, price.Ask-trans.PriceOpen)
		} else {
			return
		}
	}
}
func (src *PositionService) getProfitByBid(ch chan *model.GeneratedPrice, trans *model.Transaction) {
	for {
		price, ok := <-ch
		if ok {
			log.Printf("For position %v profit if close: %v", trans.ID, price.Bid-trans.PriceOpen)
		} else {
			return
		}
	}
}
