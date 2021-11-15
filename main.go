package main

import (
	"github.com/do1019/web-app-introduction/handler"
	"github.com/do1019/web-app-introduction/handler/middleware"
	"log"
	"net/http"
)

func main() {
	if err := mainReturnWithError(); err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func mainReturnWithError() error {
	mux := http.NewServeMux()
	mux.Handle("/do-panic", handler.NewDoPanicHandler())
	return http.ListenAndServe(":8080",
		middleware.SetDeviceOSInfoInContext(
			middleware.OutputAccessLog(
				middleware.Recovery(mux))))
}
