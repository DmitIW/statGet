package main

import (
	"log"
	"statGet/cmd/api"
	"statGet/cmd/config"
	"statGet/cmd/stop"
	"statGet/cmd/tConnector"
	"sync"
)

var (
	server api.StatisticServer
	wg     *sync.WaitGroup
)

func init() {
	if err := config.LoadEnvFile(); err != nil {
		log.Fatalf("No .env file found")
	}
	tConnector.DefConnInit(config.GetTAddr(), nil)
	server = api.GetServer(config.GetSAddr(), tConnector.GetDefaultConnection())
	wg = &sync.WaitGroup{}
	stop.Bind()
}

func main() {
	log.Println("Statistic server started")
	defer func() {
		log.Println("Statistic server stopped")
	}()
	defer wg.Wait()
	defer tConnector.DefConnClose()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.Start(); err != nil {
			log.Printf("Error on server performing: %v\n", err)
			stop.Stop()
		}
	}()
	defer server.Stop()

	stop.Wait()
}
