package handler

import (
	"net/http"
	//"fmt"
)

type DoPanicHandler struct{}

func NewDoPanicHandler() *DoPanicHandler {
	return &DoPanicHandler{}
}

func (d *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//middleware.device_os_analysis(r)
	//fmt.Println("loop")
	panic("do-panic!")
}