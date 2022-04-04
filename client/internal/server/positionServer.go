// Package server contains grpc server logic
package server

import (
	"context"
	"sync"

	"github.com/EgMeln/broker/client/internal/model"
	protocolPosition "github.com/EgMeln/broker/position_service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// PositionServer struct for grpc server logic
type PositionServer struct {
	mu           *sync.RWMutex
	generatedMap map[string]*model.GeneratedPrice
	posService   protocolPosition.PositionServiceClient
}

// ConnectPositionServer for connect to grpc server
func ConnectPositionServer() protocolPosition.PositionServiceClient {
	addressGRPC := "position_service:8083"
	con, err := grpc.Dial(addressGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolPosition.NewPositionServiceClient(con)
}

// NewPositionServer returns new service instance
func NewPositionServer(generatedMap map[string]*model.GeneratedPrice, mutex *sync.RWMutex) *PositionServer {
	return &PositionServer{
		generatedMap: generatedMap,
		mu:           mutex,
		posService:   ConnectPositionServer(),
	}
}

// OpenPositionAsk method open position record by ask
func (s *PositionServer) OpenPositionAsk(currency string) string { //nolint:dupl //Different business logic
	mod := &protocolPosition.Transaction{ID: ((s.generatedMap)[currency].ID).String(), PriceOpen: float32((s.generatedMap)[currency].Ask), IsBay: true, Symbol: currency}
	open, err := s.posService.OpenPositionAsk(context.Background(), &protocolPosition.OpenRequest{Trans: mod})
	if err != nil {
		log.Printf("Error while opening position: %v", err)
	}
	log.Printf("Position open with id: %s", open.ID)
	return open.ID
}

// OpenPositionBid method open position record by bid
func (s *PositionServer) OpenPositionBid(currency string) string { //nolint:dupl //Different business logic
	mod := &protocolPosition.Transaction{ID: ((s.generatedMap)[currency].ID).String(), PriceOpen: float32((s.generatedMap)[currency].Bid), IsBay: true, Symbol: currency}
	open, err := s.posService.OpenPositionBid(context.Background(), &protocolPosition.OpenRequest{Trans: mod})
	if err != nil {
		log.Printf("Error while opening position: %v", err)
	}
	log.Printf("Position open with id: %s", open.ID)
	return open.ID
}

// ClosePositionAsk method close position record by ask
func (s *PositionServer) ClosePositionAsk(id, currency string) {
	res, err := s.posService.ClosePositionAsk(context.Background(), &protocolPosition.CloseRequest{ID: id, Symbol: currency, PriceClose: float32((s.generatedMap)[currency].Ask)})
	if err != nil {
		log.Printf("Error while closing position: %v", err)
	} else {
		log.Printf("Position with id: %s closed. Profit %v", id, res.Result)
	}
}

// ClosePositionBid method open position record by bid
func (s *PositionServer) ClosePositionBid(id, currency string) {
	res, err := s.posService.ClosePositionBid(context.Background(), &protocolPosition.CloseRequest{ID: id, Symbol: currency, PriceClose: float32((s.generatedMap)[currency].Bid)})
	if err != nil {
		log.Printf("Error while closing position: %v", err)
	} else {
		log.Printf("Position with id: %s closed.Profit %v", id, res.Result)
	}
}
