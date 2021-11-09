package handler

import (
	"net/http"
	//"fmt"
	//"github.com/do1019/web-app-introduction/handler/middleware"
)

type DoPanicHandler struct{}

func NewDoPanicHandler() *DoPanicHandler {
	return &DoPanicHandler{}
}

func (d *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("do-panic!")
}