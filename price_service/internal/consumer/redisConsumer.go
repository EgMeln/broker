// Package consumer to consume messages from redis
package consumer

import (
	"context"
	"sync"

	"github.com/EgMeln/broker/price_service/internal/model"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// Consumer struct for redis client
type Consumer struct {
	RedisClient  *redis.Client
	redisStream  string
	mu           *sync.RWMutex
	generatedMap map[string]*model.GeneratedPrice
}

// NewConsumer returns new instance of redis consumer
func NewConsumer(ctx context.Context, cln *redis.Client, priceMap map[string]*model.GeneratedPrice, mu *sync.RWMutex) *Consumer {
	red := &Consumer{RedisClient: cln, redisStream: "STREAM", generatedMap: priceMap, mu: mu}
	go red.GetPrices(ctx)
	return red
}

// GetPrices get messages from redis
func (cons *Consumer) GetPrices(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			streams, err := cons.RedisClient.XRead(&redis.XReadArgs{
				Streams: []string{cons.redisStream, "$"},
				Count:   1,
				Block:   0,
			}).Result()
			if err != nil {
				log.Errorf("redis start process error %v", err)
			}
			if streams[0].Messages == nil {
				log.Errorf("empty message")
				continue
			}
			stream := streams[0].Messages[0].Values
			for _, value := range stream {
				price := new(model.GeneratedPrice)
				err = price.UnmarshalBinary([]byte(value.(string)))
				if err != nil {
					log.Errorf("can't parse message %v", err)
				}
				cons.mu.Lock()
				cons.generatedMap[price.Symbol] = price
				log.Info("Consumer ", cons.generatedMap[price.Symbol])
				cons.mu.Unlock()
			}
		}
	}
}
