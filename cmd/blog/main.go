package main

import (
	"context"
	"github.com/KAI-Back-end/Blog/internal/api/server"
	"github.com/KAI-Back-end/Blog/internal/config"
	"log"
	"time"
)

func main() {
	//TODO init config +
	//TODO init logger
	//TODO init server +
	//TODO init postgresql
	//TODO init traces
	//TODO init metrics
	//TODO init pprof
	//TODO init endpoints
	//TODO add closer pattern

	ctx := context.Background()

	cfg, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	srv := server.New(server.WithConfig(&cfg.Server))

	go func() {
		if err = srv.Run(ctx); err != nil {
			log.Println(err)
		}
	}()

	time.Sleep(10 * time.Minute)
}
