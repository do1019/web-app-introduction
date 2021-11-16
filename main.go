package main

import (
	//"context"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/do1019/web-app-introduction/handler"
	"github.com/do1019/web-app-introduction/handler/middleware"
	//"github.com/110y/run"
)

// func main() {
// 	if err := mainReturnWithError(); err != nil {
// 		log.Fatalln("main: failed to exit successfully, err =", err)
// 	}
// }

// func mainReturnWithError() error {
// 	mux := http.NewServeMux()

// 	ctx, stop := signal.NotifyContext(context.Background(),
// 		os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)

// 	defer stop()

// 	mux.Handle("/do-panic", handler.NewDoPanicHandler())
// 	mux.Handle("/put-count", handler.NewPutCountHandler())

// 	server := &http.Server{
// 		Addr: ":8080",
// 		Handler: middleware.ObtainIdAndPassFromEnviron().AccessRestriction(
// 			middleware.SetDeviceOSInfoInContext(
// 				middleware.OutputAccessLog(
// 					middleware.Recovery(mux)))),
// 	}

// 	errC := make(chan error, 1)

// 	go func() {
// 		<-ctx.Done()
// 		stop()
// 		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 		defer cancel()
// 		if err := server.Shutdown(ctx); err != nil {
// 			log.Println("Failed to gracefully shutdown:", err)
// 		}
// 		close(errC)
// 	}()

// 	if err := server.ListenAndServe(); err != http.ErrServerClosed {
// 		log.Println("unexpected server error", err)
// 		return err
// 	}

// 	<-errC

// 	return nil
// }

// func main() {
// 	run.Run(func(ctx context.Context) int {
// 		// Spin up your server here.
// 		mux := http.NewServeMux()
// 		mux.Handle("/put-count", handler.NewPutCountHandler())
// 		http.ListenAndServe(":8080", middleware.ObtainIdAndPassFromEnviron().AccessRestriction(
// 						middleware.SetDeviceOSInfoInContext(
// 							middleware.OutputAccessLog(
// 								middleware.Recovery(mux)))),)
// 		// This blocks until one of termination signals (unix.SIGHUP, unix.SIGINT, unix.SIGTERM or unix.SIGQUIT) will be passed.
// 		<-ctx.Done()
// 		// Tear down your server here.

// 		// After this function has finished, run.Run exits the process with returned value of this function as its exit code.
// 		return 0
// 	})
// }

func main() {
    mux := http.NewServeMux()
    mux.Handle("/put-count", handler.NewPutCountHandler())
    srv := &http.Server{
        Addr:    ":8080",
        Handler: middleware.ObtainIdAndPassFromEnviron().AccessRestriction(
			middleware.SetDeviceOSInfoInContext(
				middleware.OutputAccessLog(
					middleware.Recovery(mux)))),
    }

    go func() {
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            // Error starting or closing listener:
            log.Fatalln("Server closed with error:", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
    log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        // Error from closing listeners, or context timeout:
        log.Println("Failed to gracefully shutdown:", err)
    }
    log.Println("Server shutdown")
}