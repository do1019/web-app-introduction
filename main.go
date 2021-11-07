package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"server"
	"github.com/do1019/web-app-introduction/handler"
	"github.com/do1019/web-app-introduction/handler/middleware"
)

func main() {
	if err := mainReturnWithError(); err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func mainReturnWithError() error {
	//STEP6
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	mux := http.NewServeMux()

	//mux.Handle("/do-panic", middleware.Recovery(handler.NewDoPanicHandler()))

	//STEP3
	//mux.Handle("/do-panic", middleware.OutputAccessLog(middleware.Recovery(handler.NewDoPanicHandler())))

	//STEP4
	//mux.Handle("/do-panic", middleware.OutputAccessLog(middleware.Recovery(middleware.ObtainIdAndPassFromEnviron().AccessRestriction(handler.NewDoPanicHandler()))))

	mux.Handle("/put-count", middleware.OutputAccessLog(middleware.Recovery(middleware.ObtainIdAndPassFromEnviron().AccessRestriction(handler.NewPutCountHandler()))))
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()
	return http.ListenAndServe(":8080", mux)
}
