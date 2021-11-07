package handler

import (
	"fmt"
	"time"
	"net/http"
)

type PutCountHandler struct{}

func (p *PutCountHandler)NewPutCountHandler() *PutCountHandler {
	return &PutCountHandler{}
}

func (d *PutCountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go func() {
		for i := 10; i > 0; i--{
			fmt.Println("Until server shutdown: " + i)
			time.Sleep(2 * time.Second)
		}
	}
}