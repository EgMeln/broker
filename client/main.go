package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/EgMeln/broker/client/internal/model"
	"github.com/EgMeln/broker/client/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	mute := new(sync.RWMutex)
	priceMap := map[string]*model.GeneratedPrice{
		"Aeroflot": {},
		"ALROSA":   {},
		"Akron":    {},
	}
	ctx := context.Background()
	priceClient := server.ConnectPriceServer()
	log.Infof("start")
	go server.SubscribePrices(ctx, "Aeroflot", priceClient, mute, priceMap)
	go server.SubscribePrices(ctx, "ALROSA", priceClient, mute, priceMap)
	go server.SubscribePrices(ctx, "Akron", priceClient, mute, priceMap)
	posClient := server.NewPositionServer(priceMap, mute)

	log.Infof("Start open")
	id2 := posClient.OpenPositionAsk(ctx, "ALROSA")
	id := posClient.OpenPositionBid(ctx, "Aeroflot")
	time.Sleep(40 * time.Second)
	posClient.ClosePositionAsk(ctx, id2, "ALROSA")
	posClient.ClosePositionBid(ctx, id, "Aeroflot")
	//t := time.Now()
	//var array []string
	//for i := 0; i < 10000; i++ {
	//	id := posClient.OpenPositionAsk(ctx, "Aeroflot")
	//	array = append(array, id)
	//	//time.Sleep(50 * time.Millisecond)
	//}
	//time.Sleep(10 * time.Second)
	//for _, id := range array {
	//	posClient.ClosePositionAsk(ctx, id, "Aeroflot")
	//}
	//log.Info(time.Since(t))
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-c
	log.Info("END")
}
