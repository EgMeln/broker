package main

import (
	"context"
	"fmt"
	"github.com/EgMeln/broker/position_service/internal/config"
	"github.com/EgMeln/broker/position_service/internal/model"
	"github.com/EgMeln/broker/position_service/internal/repository"
	"github.com/EgMeln/broker/position_service/internal/server"
	"github.com/EgMeln/broker/position_service/internal/service"
	"github.com/EgMeln/broker/position_service/protocol"
	"github.com/EgMeln/broker/price_service"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"sync"
)

func main() {
	initLog()
	cfg, err := config.New()
	if err != nil {
		log.Warnf("Config error %v", err)
	}

	cfg.DBURL = fmt.Sprintf("%s://%s:%s@%s:%d/%s", cfg.DB, cfg.User, cfg.Password, cfg.Host, cfg.PortPostgres, cfg.DBNamePostgres)
	log.Infof("DB URL: %s", cfg.DBURL)
	pool := connectPostgres(cfg.DBURL)
	log.Infof("Connected!")

	mu := new(sync.RWMutex)

	transactionMap := map[string]*model.GeneratedPrice{}

	transactionService := service.NewPositionService(&repository.PostgresPrice{PoolPrice: pool})

	transactionServer := server.NewPositionServer(*transactionService, mu, &transactionMap)

	err = runGRPC(transactionServer)
	if err != nil {
		log.Printf("err in grpc run %v", err)
	}
}
func initLog() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func connectPostgres(URL string) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), URL)
	if err != nil {
		log.Warnf("Error connection to DB %v", err)
	}
	return pool
}
func runGRPC(recServer protocol.PositionServiceServer) error {
	port := "localhost:8083"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protocol.RegisterPositionServiceServer(grpcServer, recServer)
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return grpcServer.Serve(listener)
}

//func subscribePrices(symbol string, client protocol., mu *sync.RWMutex, transactionMap map[string]*model.GeneratedPrice) {
//	req := protocol.OpenRequest{Trans: }
//}
