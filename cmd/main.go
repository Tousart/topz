package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/tousart/topz/internal/api"
	"github.com/tousart/topz/internal/server"
	"github.com/tousart/topz/internal/service"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	wg := &sync.WaitGroup{}

	errChan := make(chan error, 1)

	procService := service.NewProcService()

	topzApi := api.NewTopzApi(procService)
	mux := api.NewMux()
	topzApi.WithHandlers(mux)

	// default port=9000
	server.CreateAndrunServer(ctx, mux, fmt.Sprintf(":%s", os.Getenv("PORT")), errChan, wg)

	wg.Wait()
}
