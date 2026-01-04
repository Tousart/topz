package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/tousart/topz/internal/api"
)

func CreateAndrunServer(ctx context.Context, mux *api.Mux, addr string, errChan chan error, wg *sync.WaitGroup) {
	serv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		log.Printf("topz run on %s\n", addr)
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
			fmt.Fprintf(os.Stderr, "failed to run topz on %s, error: %v", addr, err)
			return
		}
		fmt.Println("topz stopped")
	}()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg.Go(func() {
		select {
		case <-ctx.Done():
		case <-errChan:
		}

		if err := serv.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down topz: %v", err)
			return
		}
		fmt.Println("topz shutting down graceful")
	})
}
