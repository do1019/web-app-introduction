package main

import (
	"log"
	"net/http"
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
	mux.Handle("/do-panic", &handler.DoPanicHandler{})
	return http.ListenAndServe(":8080", middleware.Recovery(mux))
}
