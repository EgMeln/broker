package repository

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/EgMeln/broker/position_service/internal/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// OpenPosition function to inserting a position into postgres db
func (rep *PostgresPrice) OpenPosition(ctx context.Context, trans *model.Transaction, bay string) (*uuid.UUID, error) {
	_, err := rep.PoolPrice.Exec(ctx, "INSERT INTO positions(id_,price_open,is_bay,symbol,price_close,bay_by) VALUES ($1,$2,$3,$4,$5,$6)",
		trans.ID, trans.PriceOpen, true, trans.Symbol, trans.PriceClose, bay)
	if err != nil {
		log.Errorf("can't insert position %v", err)
		return &trans.ID, err
	}
	return &trans.ID, nil
}

// ClosePosition function to closing a position into postgres db
func (rep *PostgresPrice) ClosePosition(ctx context.Context, closePrice *float64, id *uuid.UUID) (string, error) {
	row, err := rep.PoolPrice.Exec(ctx, "UPDATE positions SET price_close = $1,is_bay = $2 WHERE id_ = $3",
		closePrice, false, id)
	if row.RowsAffected() == 0 {
		log.Errorf("rows empty %v", err)
		return "", exec.ErrNotFound
	}
	if err != nil {
		log.Errorf("can't update position %v", err)
		return "", err
	}
	var openPrice *float64
	err = rep.PoolPrice.QueryRow(ctx, "SELECT price_open from positions where id_=$1", id).Scan(&openPrice)
	str := fmt.Sprintf("Position close: profit: %v", *closePrice-*openPrice)
	return str, err
}
