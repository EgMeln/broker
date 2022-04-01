package main

import (
	"github.com/EgMeln/broker/client/internal/model"
	"github.com/EgMeln/broker/client/internal/server"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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

	posClient := server.NewPositionServer(&priceMap, mute)

	log.Infof("Start open")
	id := posClient.OpenPositionAsk("Aeroflot")

	time.Sleep(5 * time.Second)

	log.Infof("Start close")
	posClient.ClosePositionAsk(id, "Aeroflot")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-c
	log.Info("END")
}
