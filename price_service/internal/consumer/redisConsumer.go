package consumer

import (
	"context"
	"github.com/EgMeln/broker/price_service/internal/model"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Consumer struct {
	RedisClient  *redis.Client
	redisStream  string
	mu           sync.RWMutex
	generatedMap *map[string]*model.GeneratedPrice
}

func NewConsumer(ctx context.Context, cln *redis.Client, priceMap *map[string]*model.GeneratedPrice) *Consumer {
	red := &Consumer{RedisClient: cln, redisStream: "STREAM", generatedMap: priceMap, mu: sync.RWMutex{}}
	go red.GetPrices(ctx)
	return red
}
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
			price := new(model.GeneratedPrice)
			for _, value := range stream {
				err = price.UnmarshalBinary([]byte(value.(string)))
				if err != nil {
					log.Errorf("can't parse message %v", err)
				}
				cons.mu.Lock()
				(*cons.generatedMap)[price.Symbol] = price
				log.Info((*cons.generatedMap)[price.Symbol])
				cons.mu.Unlock()
			}
		}
	}
}
