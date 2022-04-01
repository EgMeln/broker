package service

import (
	"context"
	"github.com/EgMeln/broker/position_service/internal/model"
	"github.com/EgMeln/broker/position_service/internal/repository"
	"github.com/google/uuid"
)

type PositionService struct {
	rep repository.PriceTransaction
}

func NewPositionService(rep *repository.PostgresPrice) *PositionService {
	return &PositionService{rep: rep}
}

func (src *PositionService) OpenPosition(ctx context.Context, trans *model.Transaction) (*uuid.UUID, error) {
	return src.rep.OpenPosition(ctx, trans)
}

func (src *PositionService) ClosePosition(ctx context.Context, closePrice *float64, id *uuid.UUID) (string, error) {
	return src.rep.ClosePosition(ctx, closePrice, id)
}
