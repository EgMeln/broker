package server

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/EgMeln/broker/client/internal/model"
	protocolPrice "github.com/EgMeln/broker/price_service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ConnectPriceServer for connect to grpc server
func ConnectPriceServer() protocolPrice.PriceServiceClient {
	addressGRPC := "price_service:8089"
	con, err := grpc.Dial(addressGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolPrice.NewPriceServiceClient(con)
}

// SubscribePrices for get prices from grpc
func SubscribePrices(ctx context.Context, symbol string, client protocolPrice.PriceServiceClient, mu *sync.RWMutex, transactionMap map[string]*model.GeneratedPrice) {
	req := protocolPrice.GetRequest{Symbol: symbol}
	stream, err := client.GetPrice(ctx, &req)
	i := 0
	t := time.Now()
	if err != nil {
		log.Fatalf("%v get price error, %v", client, err)
	}
	for {
		if i == 10 {
			i = 0
			log.Info(time.Since(t))
			t = time.Now()
		}
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a note : %v", err)
		}

		cur := &model.GeneratedPrice{Symbol: in.Price.Symbol, Ask: float64(in.Price.Ask), Bid: float64(in.Price.Bid), DoteTime: in.Price.Time}
		mu.Lock()
		transactionMap[cur.Symbol] = cur
		mu.Unlock()

		//log.Infof("Got currency data Name: %v Ask: %v Bid: %v  at time %v", in.Price.Symbol, in.Price.Ask, in.Price.Bid, in.Price.Time)
		i++

	}
}
