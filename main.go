package main

import (
	"net/http"

	"github.com/do1019/web-app-introduction/handler"
	"github.com/do1019/web-app-introduction/handler/middleware"
)

func main() {
	mux := http.NewServeMux()

	//mux.Handle("/do-panic", middleware.Recovery(handler.NewDoPanicHandler()))

	//STEP3
	//mux.Handle("/do-panic", middleware.OutputAccessLog(middleware.Recovery(handler.NewDoPanicHandler())))

	//STEP4
	mux.Handle("/do-panic", middleware.OutputAccessLog(middleware.Recovery(middleware.ObtainIdAndPassFromEnviron().AccessRestriction(handler.NewDoPanicHandler()))))
	// ap := middleware.ObtainIdAndPassFromEnviron()
	// mux.Handle("/do-panic", middleware.OutputAccessLog(middleware.Recovery(ap.AccessRestriction(handler.NewDoPanicHandler()))))
	http.ListenAndServe(":8080", mux)
}
