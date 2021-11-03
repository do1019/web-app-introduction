package handler

import (
	"net/http"
	//"fmt"
	"github.com/do1019/web-app-introduction/handler/middleware"
)

type DoPanicHandler struct{}

func NewDoPanicHandler() *DoPanicHandler {
	return &DoPanicHandler{}
}

func (d *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware.Device_os_analysis(r)
	//fmt.Println("loop")
	panic("do-panic!")
}