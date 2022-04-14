// Package service contains business logic
package service

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"sync"

	"github.com/EgMeln/broker/position_service/internal/model"
	"github.com/EgMeln/broker/position_service/internal/repository"
	"github.com/google/uuid"
)

// PositionService struct for
type PositionService struct {
	pool           *pgxpool.Pool
	rep            repository.PriceTransaction
	mu             *sync.RWMutex
	positionMap    map[string]map[string]*chan *model.GeneratedPrice
	transactionMap map[string]*model.Transaction
}

// NewPositionService used for setting position services
func NewPositionService(ctx context.Context, rep *repository.PostgresPrice, pos map[string]map[string]*chan *model.GeneratedPrice, pool *pgxpool.Pool, mute *sync.RWMutex) *PositionService {
	PosService := PositionService{rep: rep, mu: mute, positionMap: pos, pool: pool, transactionMap: make(map[string]*model.Transaction)}
	go PosService.waitForNotification(ctx)
	return &PosService
}

// OpenPosition add record about position
func (src *PositionService) OpenPosition(ctx context.Context, trans *model.Transaction, str string) (*uuid.UUID, error) {
	return src.rep.OpenPosition(ctx, trans, str)
}

// ClosePosition update record about position
func (src *PositionService) ClosePosition(ctx context.Context, closePrice *float64, id *uuid.UUID) (string, error) {
	return src.rep.ClosePosition(ctx, closePrice, id)
}

func (src *PositionService) getProfitByAsk(ctx context.Context, ch chan *model.GeneratedPrice, trans *model.Transaction) {
	for {
		price, ok := <-ch
		if ok {
			log.Infof("For position %v profit if close: %v", trans.ID, price.Ask-trans.PriceOpen)
			go src.SystemStop(ctx, price, trans)
		} else {
			log.Infof("Position with id %v close", trans.ID)
			return
		}
	}
}
func (src *PositionService) getProfitByBid(ctx context.Context, ch chan *model.GeneratedPrice, trans *model.Transaction) {
	for {
		price, ok := <-ch
		if ok {
			log.Infof("For position %v profit if close: %v", trans.ID, price.Bid-trans.PriceOpen)
			go src.SystemStop(ctx, price, trans)
		} else {
			log.Infof("Position with id %v close", trans.ID)
			return
		}
	}
}
func (src *PositionService) waitForNotification(ctx context.Context) {
	conn, err := src.pool.Acquire(ctx)
	if err != nil {
		log.Errorf("Error connection %v", err)
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "listen positions")
	if err != nil {
		log.Errorf(" conn exec %v", err)
	}
	for {
		notification, err := conn.Conn().WaitForNotification(ctx)
		if err != nil {
			log.Errorf("error waiting for notification: %v", err)
		}
		position := model.Transaction{}
		if err := json.Unmarshal([]byte(notification.Payload), &position); err != nil {
			log.Errorf("Unmarshal error %v", err)
		}
		if err != nil {
			log.Error("Error waiting for notification:", err)
		}
		ch := make(chan *model.GeneratedPrice)
		if position.IsBay {
			src.mu.Lock()
			src.positionMap[position.Symbol][position.ID.String()] = &ch
			src.transactionMap[position.ID.String()] = &position
			src.mu.Unlock()
			if position.BayBy == "Ask" {
				go src.getProfitByAsk(ctx, ch, &position)
			} else if position.BayBy == "Bid" {
				go src.getProfitByBid(ctx, ch, &position)
			}
		} else {
			src.mu.Lock()
			close(*src.positionMap[position.Symbol][position.ID.String()])
			delete(src.positionMap[position.Symbol], position.ID.String())
			delete(src.transactionMap, position.ID.String())
			src.mu.Unlock()
		}
	}
}
func (src *PositionService) SystemStop(ctx context.Context, update *model.GeneratedPrice, transaction *model.Transaction) {
	src.mu.RLock()
	if transaction.Symbol == update.Symbol {
		if stopLoss(update, transaction) || takeProfit(update, transaction) {
			var stopPrice float64
			if transaction.BayBy == "Ask" {
				stopPrice = update.Ask
			} else if transaction.BayBy == "Bid" {
				stopPrice = update.Bid
			}
			profit, err := src.ClosePosition(ctx, &stopPrice, &transaction.ID)
			if err != nil {
				log.Errorf("close positins error")
			}
			log.Info("profit ", profit)
		}
	}
	src.mu.RUnlock()
	src.mu.Lock()
	positionID, priceClose, ifClose := src.marginLiquidation(update)
	if ifClose {
		profit, err := src.ClosePosition(ctx, &priceClose, &positionID)
		if err != nil {
			log.Errorf("close positins error")
		}
		log.Info("profit ", profit)
	}
	src.mu.Unlock()
}
func (src *PositionService) marginLiquidation(pos *model.GeneratedPrice) (uuid.UUID, float64, bool) {
	var balance float64
	var positionID uuid.UUID
	var priceClose float64
	for _, position := range src.transactionMap {
		if position.Symbol == pos.Symbol {
			if position.BayBy == "Ask" {
				balance += position.PriceOpen - pos.Ask
				if (position.PriceOpen - pos.Ask) <= 0 {
					positionID = position.ID
					priceClose = pos.Ask
				}
			} else if position.BayBy == "Bid" {
				balance += position.PriceOpen - pos.Bid
				if (position.PriceOpen - pos.Bid) <= 0 {
					positionID = position.ID
					priceClose = pos.Bid
				}
			}
		}
	}
	return positionID, priceClose, balance < 0.0
}
func takeProfit(pos *model.GeneratedPrice, trans *model.Transaction) bool {
	if trans.IsBay {
		if trans.BayBy == "Ask" {
			return trans.TakeProfit <= pos.Ask
		} else if trans.BayBy == "Bid" {
			return trans.TakeProfit <= pos.Bid
		}
	}
	return false
}
func stopLoss(pos *model.GeneratedPrice, trans *model.Transaction) bool {
	if trans.IsBay {
		if trans.BayBy == "Ask" {
			return trans.StopLoss >= pos.Ask
		} else if trans.BayBy == "Bid" {
			return trans.StopLoss >= pos.Bid
		}
	}
	return false
}
