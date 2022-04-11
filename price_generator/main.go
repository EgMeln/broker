package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/EgMeln/broker/price_generator/internal/config"
	"github.com/EgMeln/broker/price_generator/internal/producer"
	"github.com/EgMeln/broker/price_generator/internal/service/generate"
	"github.com/EgMeln/broker/price_generator/internal/service/send"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

func main() {
	redisCfg, err := config.NewRedis()
	if err != nil {
		log.Fatalln("Config error: ", redisCfg)
	}
	ctx := context.Background()
	redisClient := connRedis(ctx, redisCfg)
	gen := generate.NewGenerator()
	prod := producer.NewRedis(redisClient)
	serv := send.NewService(prod, gen)

	log.Info(gen)

	log.Info("Start generating prices")

	ctx, cancel := context.WithCancel(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	err = serv.StartSending(ctx)
	if err != nil {
		log.Fatalf("Sending error %v", err)
	}

	log.Println("received signal", <-c)
	cancel()
	err = redisClient.Close()
	if err != nil {
		log.Fatalf("redis close error %v", err)
	}
	log.Info("Success consuming messages")
}

func connRedis(ctx context.Context, redisCfg *config.RedisConfig) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB})
	if _, ok := redisClient.Ping(ctx).Result(); ok != nil {
		log.Fatalf("redis new client error %v", ok)
	}
	return redisClient
}
