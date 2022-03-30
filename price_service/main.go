package main

import (
	"context"
	"github.com/EgMeln/broker/price_service/internal/config"
	"github.com/EgMeln/broker/price_service/internal/consumer"
	"github.com/EgMeln/broker/price_service/internal/model"
	"github.com/EgMeln/broker/price_service/internal/server"
	"github.com/EgMeln/broker/price_service/protocol"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	redisCfg, err := config.NewRedis()
	if err != nil {
		log.Fatalln("Config error: ", redisCfg)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB})

	priceMap := map[string]*model.GeneratedPrice{
		"Aeroflot": {},
		"ALROSA":   {},
		"Akron":    {},
	}

	mutex := sync.RWMutex{}

	priceServer := server.NewPriceServer(&mutex, &priceMap)

	go runGRPC(priceServer)

	ctx, cancel := context.WithCancel(context.Background())
	cons := consumer.NewConsumer(ctx, redisClient, &priceMap)

	defer func(RedisClient *redis.Client) {
		err := RedisClient.Close()
		if err != nil {
			log.Fatalln("close redis connection error %v", err)
		}
	}(cons.RedisClient)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-c
	cancel()
}

func runGRPC(priceServer *server.PriceServer) {
	listener, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protocol.RegisterPriceServiceServer(grpcServer, priceServer)
	log.Printf("server listening at %v", listener.Addr())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("connect grpc error %v", err)
	}
}
