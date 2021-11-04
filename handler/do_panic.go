package handler

import (
	"net/http"
	"fmt"
	"github.com/do1019/web-app-introduction/handler/middleware"
)

type DoPanicHandler struct{}

func NewDoPanicHandler() *DoPanicHandler {
	return &DoPanicHandler{}
}

func (d *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// if v := ctx.Value("OS"); v != nil {
	// 	fmt.Println("found value:", v)
	// }
	// fmt.Println("key not found:", "OS")
	// nr := middleware.SetDeviceOSInfoInContext(r)
	// v, err := middleware.GetDeviceOSInfoInContext(nr)
	// if err != nil {
	// 	fmt.Println("key not found:", "OS", err)
	// }
	// fmt.Println("found value:", v)
	
	panic("do-panic!")
}