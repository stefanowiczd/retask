package main

import (
	"context"
	"errors"
	"github.com/stefanowiczd/retask/internal/application"
	"github.com/stefanowiczd/retask/internal/interface/rest"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	handlerPackageManager := rest.NewHandlerPacksManager(
		application.NewServicePacksManager(),
	)

	svr := rest.NewServer(
		rest.DefaultConfig(),
		handlerPackageManager,
	)

	run(svr)
}

func run(s *rest.Server) {
	waitGroup := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()

		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		case sig := <-signalChan:
			log.Printf("Received signal: %s\n", sig)
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server failed: %v\n", err)
			cancel()
		}
	}()

	<-ctx.Done()

	shutDownCtx, shutDownCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer shutDownCancel()

	if err := s.Shutdown(shutDownCtx); err != nil {
		log.Printf("Unable graceful server shutdown: %v\n", err)
	}

	waitGroup.Wait()
}
