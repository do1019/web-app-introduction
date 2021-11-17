package main

import (
	"context"
	"github.com/do1019/web-app-introduction/handler"
	"github.com/do1019/web-app-introduction/handler/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	if err := mainReturnWithError(); err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func mainReturnWithError() error {
	mux := http.NewServeMux()

	mux.Handle("/do-panic", handler.NewDoPanicHandler())
	mux.Handle("/put-count", handler.NewPutCountHandler())

	server := &http.Server{
		Addr: ":8080",
		Handler: middleware.ObtainIdAndPassFromEnviron().AccessRestriction(
			middleware.SetDeviceOSInfoInContext(
				middleware.OutputAccessLog(
					middleware.Recovery(mux)))),
	}

	return GracefulShutdown(server)
}

func GracefulShutdown(server *http.Server) error {
	wg := sync.WaitGroup{}
	ctx, _ := signal.NotifyContext(context.Background(),
		os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-ctx.Done()
		wg.Add(1)
		log.Printf("SIGNAL %d received, then shutting down...\n", <-ctx.Done())
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Println("Failed to gracefully shutdown:", err)
		}
		wg.Done()
	}()
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Println("unexpected server error", err)
		return err
	}
	wg.Wait()
	log.Println("Server shutdown")
	return nil
}
