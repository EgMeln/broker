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
	positionMap  map[string]map[string]*model.Transaction
}

// NewPositionService used for setting position services
func NewPositionService(rep *repository.PostgresPrice, priceMap map[string]*model.GeneratedPrice, mute *sync.RWMutex, pos map[string]map[string]*model.Transaction) *PositionService {
	return &PositionService{rep: rep, generatedMap: priceMap, mu: mute, positionMap: pos}
}

// OpenPosition add record about position
func (src *PositionService) OpenPosition(ctx context.Context, trans *model.Transaction, str string) (*uuid.UUID, error) {
	src.mu.Lock()
	src.positionMap[str][trans.ID.String()] = trans
	src.mu.Unlock()
	if str == "Ask" {
		go src.getProfitByAsk(src.positionMap, trans)
	} else if str == "Bid" {
		go src.getProfitByBid(src.positionMap, trans)
	}
	return src.rep.OpenPosition(ctx, trans)
}

// ClosePosition update record about position
func (src *PositionService) ClosePosition(ctx context.Context, closePrice *float64, id *uuid.UUID, str string) (string, error) {
	src.positionMap[str][id.String()].IsBay = false
	return src.rep.ClosePosition(ctx, closePrice, id)
}

func (src *PositionService) getProfitByAsk(posMap map[string]map[string]*model.Transaction, trans *model.Transaction) {
	for {
		src.mu.RLock()
		if posMap["Ask"][trans.ID.String()].IsBay == true {
			log.Printf("For position %v profit if close: %v", trans.ID, src.generatedMap[trans.Symbol].Ask-trans.PriceOpen)
		} else {
			log.Printf("Position with id %v close", trans.ID)
			delete(src.positionMap["Ask"], trans.ID.String())
			return
		}
		src.mu.RUnlock()
	}
}
func (src *PositionService) getProfitByBid(posMap map[string]map[string]*model.Transaction, trans *model.Transaction) {
	for {
		if posMap["Bid"][trans.ID.String()].IsBay == true {
			src.mu.RLock()
			log.Printf("For position %v profit if close: %v", trans.ID, src.generatedMap[trans.Symbol].Bid-trans.PriceOpen)
			src.mu.RUnlock()
		} else {
			log.Printf("Position with id %v close", trans.ID)
			return
		}
	}
}
