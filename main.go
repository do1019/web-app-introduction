package main

import (
	//"fmt"
	"log"
	"net/http"
	"context"
	"os"
	"os/signal"
	"time"
	"syscall"
	"github.com/do1019/web-app-introduction/handler"
	"github.com/do1019/web-app-introduction/handler/middleware"
)

func main() {
	if err := mainReturnWithError(); err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func mainReturnWithError() error {
	mux := http.NewServeMux()

	//STEP3
	//mux.Handle("/do-panic", handler.NewDoPanicHandler())
	//http.ListenAndServe(":8080", middleware.OutputAccessLog(middleware.Recovery(mux)))

	//STEP4
	//mux.Handle("/do-panic", handler.NewDoPanicHandler())
	//http.ListenAndServe(":8080", middleware.OutputAccessLog(middleware.Recovery(middleware.ObtainIdAndPassFromEnviron().AccessRestriction(mux))))

	//STEP6
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)

	defer stop()

	mux.Handle("/put-count", middleware.OutputAccessLog(middleware.Recovery(middleware.ObtainIdAndPassFromEnviron().AccessRestriction(handler.NewPutCountHandler()))))
	//mux.Handle("/put-count", handler.NewPutCountHandler())

	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.OutputAccessLog(middleware.Recovery(middleware.ObtainIdAndPassFromEnviron().AccessRestriction(handler.NewPutCountHandler()))),
	}

	errC := make(chan error, 1)

	go func() {
		<-ctx.Done()
		stop()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Println("Failed to gracefully shutdown:", err)
		}
		close(errC)
	}()
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Println("unexpected server error", err)
		return err
	}
	<-errC
	return nil
}
