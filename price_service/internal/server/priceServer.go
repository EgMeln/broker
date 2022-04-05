// Package server contains grpc server logic
package server

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"

	"github.com/EgMeln/broker/price_service/internal/model"
	"github.com/EgMeln/broker/price_service/protocol"
)

// PriceServer struct for grpc server logic
type PriceServer struct {
	mu           *sync.RWMutex
	generatedMap map[string]*model.GeneratedPrice
	protocol.UnimplementedPriceServiceServer
}

// NewPriceServer returns new service instance
func NewPriceServer(mu *sync.RWMutex, priceMap map[string]*model.GeneratedPrice) *PriceServer {
	return &PriceServer{generatedMap: priceMap, mu: mu}
}

// GetPrice method get prices from redis stream
func (priceServ *PriceServer) GetPrice(in *protocol.GetRequest, stream protocol.PriceService_GetPriceServer) error {
	key := in.Symbol
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			priceServ.mu.RLock()
			resp := priceServ.generatedMap[key]
			priceServ.mu.RUnlock()
			cur := protocol.Price{Symbol: resp.Symbol, Ask: float32(resp.Ask), Bid: float32(resp.Bid), ID: resp.ID.String(), Time: resp.DoteTime}
			err := stream.Send(&protocol.GetResponse{Price: &cur})
			log.Info("Server ", &cur)
			if err != nil {
				return err
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
