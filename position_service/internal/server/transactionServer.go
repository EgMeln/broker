package server

import (
	"context"
	"github.com/EgMeln/broker/position_service/internal/model"
	"github.com/EgMeln/broker/position_service/internal/service"
	"github.com/EgMeln/broker/position_service/protocol"
	"github.com/google/uuid"
	"sync"
)

type PositionServer struct {
	mu           *sync.RWMutex
	generatedMap *map[string]*model.GeneratedPrice
	posService   service.PositionService
	*protocol.UnimplementedPositionServiceServer
}

func NewPositionServer(serv service.PositionService, mu *sync.RWMutex, priceMap *map[string]*model.GeneratedPrice) *PositionServer {
	return &PositionServer{generatedMap: priceMap, mu: mu, posService: serv}
}

func (srv *PositionServer) OpenPositionAsk(ctx context.Context, in *protocol.OpenRequest) (*protocol.OpenResponse, error) {
	position := model.Transaction{
		ID:        uuid.New(),
		PriceOpen: (*srv.generatedMap)[in.Trans.Symbol].Ask,
		IsBay:     true,
		Symbol:    in.Trans.Symbol,
	}
	id, err := srv.posService.OpenPosition(ctx, &position)
	if err != nil {
		return nil, err
	}
	return &protocol.OpenResponse{ID: id.String()}, nil
}

func (srv *PositionServer) OpenPositionBid(ctx context.Context, in *protocol.OpenRequest) (*protocol.OpenResponse, error) {
	position := model.Transaction{
		ID:        uuid.New(),
		PriceOpen: (*srv.generatedMap)[in.Trans.Symbol].Bid,
		IsBay:     true,
		Symbol:    in.Trans.Symbol,
	}
	id, err := srv.posService.OpenPosition(ctx, &position)
	if err != nil {
		return nil, err
	}
	return &protocol.OpenResponse{ID: id.String()}, nil
}
func (srv *PositionServer) ClosePositionAsk(ctx context.Context, in *protocol.CloseRequest) (*protocol.CloseResponse, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return &protocol.CloseResponse{}, err
	}
	err = srv.posService.ClosePosition(ctx, &(*srv.generatedMap)[in.Symbol].Ask, &id)
	if err != nil {
		return &protocol.CloseResponse{}, err
	}
	return &protocol.CloseResponse{}, nil
}
func (srv *PositionServer) ClosePositionBid(ctx context.Context, in *protocol.CloseRequest) (*protocol.CloseResponse, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return &protocol.CloseResponse{}, err
	}
	err = srv.posService.ClosePosition(ctx, &(*srv.generatedMap)[in.Symbol].Bid, &id)
	if err != nil {
		return &protocol.CloseResponse{}, err
	}
	return &protocol.CloseResponse{}, nil
}
