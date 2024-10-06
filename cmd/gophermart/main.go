package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/NikolayStrekalov/practicum-gophermart/internal/infra/config"
	"github.com/NikolayStrekalov/practicum-gophermart/internal/infra/db"
	"github.com/NikolayStrekalov/practicum-gophermart/internal/infra/server"
)

func main() {
	config.InitConfig()
	err := db.InitDB(config.AppConfig.Database)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	server.Start(ctx)
	// TODO worker start

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()
}
