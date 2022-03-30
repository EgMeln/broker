package producer

import (
	"context"
	"fmt"
	"github.com/EgMeln/broker/price_generator/internal/service/generate"
	"github.com/go-redis/redis/v8"
)

type Producer struct {
	redisClient *redis.Client
	redisStream string
}

// NewRedis returns new instance of Producer
func NewRedis(cln *redis.Client) *Producer {
	return &Producer{redisClient: cln, redisStream: "STREAM"}
}

// ProduceMessage send messages to redis stream
func (cln *Producer) ProduceMessage(ctx context.Context, prices *generate.Generator) error {
	_, err := cln.redisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: cln.redisStream,
		Values: prices.Prices,
	}).Result()
	if err != nil {
		return fmt.Errorf("redis send message error %v", err)
	}
	return nil
}
