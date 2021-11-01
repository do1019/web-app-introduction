package main

import (
	"net/http"
	"github.com/do1019/web-app-introduction/handler"
	"github.com/do1019/web-app-introduction/handler/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/do-panic", middleware.Recovery(handler.NewDoPanicHandler()))
	http.ListenAndServe(":8080", mux)
}