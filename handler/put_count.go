package handler

import (
	"fmt"
	"time"
	"net/http"
	"strconv"
)

type PutCountHandler struct{}

func NewPutCountHandler() *PutCountHandler {
	return &PutCountHandler{}
}

func (p *PutCountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for i := 5; i > 0; i--{
		fmt.Println("Until server shutdown: " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}