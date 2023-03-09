package main

import (
	"github.com/JamesTiberiusKirk/subs-only/stripe-events-handler/server"
)

func main() {
	initLogger()

	config := buildConfig()

	// db := initDB(config)
	e := initHttpServer()

	server.NewServer(e)

	e.Logger.Fatal(e.Start(config.HTTP.Port))
}
