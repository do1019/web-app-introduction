package handler

import (
	"net/http"
)

type DoPanicHandler struct{}

func NewDoPanicHandler() *DoPanicHandler {
	return &DoPanicHandler{}
}

func (d *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("do-panic!")
}