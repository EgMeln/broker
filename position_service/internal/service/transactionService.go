// Package service contains business logic
package service

import (
	"context"

	"github.com/EgMeln/broker/position_service/internal/model"
	"github.com/EgMeln/broker/position_service/internal/repository"
	"github.com/google/uuid"
)

// PositionService struct for
type PositionService struct {
	rep repository.PriceTransaction
}

// NewPositionService used for setting position services
func NewPositionService(rep *repository.PostgresPrice) *PositionService {
	return &PositionService{rep: rep}
}

// OpenPosition add record about position
func (src *PositionService) OpenPosition(ctx context.Context, trans *model.Transaction) (*uuid.UUID, error) {
	return src.rep.OpenPosition(ctx, trans)
}

// ClosePosition update record about position
func (src *PositionService) ClosePosition(ctx context.Context, closePrice *float64, id *uuid.UUID) (string, error) {
	return src.rep.ClosePosition(ctx, closePrice, id)
}
