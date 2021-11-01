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
	//fmt.Println("loop")
	panic("do-panic!")
}