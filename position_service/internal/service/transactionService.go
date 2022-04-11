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
	pool         *pgxpool.Pool
	rep          repository.PriceTransaction
	generatedMap map[string]*model.GeneratedPrice
	mu           *sync.RWMutex
	positionMap  map[string]map[string]*chan *model.GeneratedPrice
}

// NewPositionService used for setting position services
func NewPositionService(ctx context.Context, rep *repository.PostgresPrice, priceMap map[string]*model.GeneratedPrice, mute *sync.RWMutex, pos map[string]map[string]*chan *model.GeneratedPrice, pool *pgxpool.Pool) *PositionService {
	PosService := PositionService{rep: rep, generatedMap: priceMap, mu: mute, positionMap: pos, pool: pool}
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

func (src *PositionService) getProfitByAsk(ch chan *model.GeneratedPrice, trans *model.Transaction) {
	for {
		price, ok := <-ch
		if ok {
			log.Printf("For position %v profit if close: %v", trans.ID, price.Ask-trans.PriceOpen)
		} else {
			log.Printf("Position with id %v close", trans.ID)
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
			log.Printf("Position with id %v close", trans.ID)
			return
		}
	}
}
func (src *PositionService) waitForNotification(ctx context.Context) {
	conn, err := src.pool.Acquire(ctx)
	if err != nil {
		log.Printf("Error connection %v", err)
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "listen positions")
	if err != nil {
		log.Printf(" conn exec %v", err)
	}
	for {
		notification, err := conn.Conn().WaitForNotification(ctx)
		if err != nil {
			log.Printf("error waiting for notification: %v", err)
		}
		position := model.Transaction{}
		if err := json.Unmarshal([]byte(notification.Payload), &position); err != nil {
			log.Printf("Unmarshal error %v", err)
		}
		if err != nil {
			log.Println("Error waiting for notification:", err)
		}
		ch := make(chan *model.GeneratedPrice)
		switch position.IsBay {
		case true:
			src.mu.Lock()
			src.positionMap[position.Symbol][position.ID.String()] = &ch
			src.mu.Unlock()
			if position.BayBy == "Ask" {
				go src.getProfitByAsk(ch, &position)
			} else if position.BayBy == "Bid" {
				go src.getProfitByBid(ch, &position)
			}
		case false:
			src.mu.Lock()
			close(*src.positionMap[position.Symbol][position.ID.String()])
			delete(src.positionMap[position.Symbol], position.ID.String())
			src.mu.Unlock()
		}
	}
}
