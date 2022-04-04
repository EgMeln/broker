package main

import (
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
	priceClient := server.ConnectPriceServer()
	log.Infof("start")
	go server.SubscribePrices("Aeroflot", priceClient, mute, priceMap)
	go server.SubscribePrices("ALROSA", priceClient, mute, priceMap)
	posClient := server.NewPositionServer(priceMap, mute)

	log.Infof("Start open")
	id := posClient.OpenPositionAsk("Aeroflot")
	id2 := posClient.OpenPositionBid("Aeroflot")
	id3 := posClient.OpenPositionAsk("ALROSA")
	id4 := posClient.OpenPositionBid("ALROSA")
	time.Sleep(20 * time.Second)

	log.Infof("Start close1")
	posClient.ClosePositionAsk(id, "Aeroflot")
	log.Infof("Start close2")
	posClient.ClosePositionBid(id2, "Aeroflot")
	log.Infof("Start close3")
	posClient.ClosePositionAsk(id3, "ALROSA")
	log.Infof("Start close4")
	posClient.ClosePositionBid(id4, "ALROSA")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-c
	log.Info("END")
}
