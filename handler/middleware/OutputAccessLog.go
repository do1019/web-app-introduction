package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func OutputAccessLog(next http.Handler) http.Handler {
	ap: := &model.Access{}
	t := time.Now()
	// ここでhandler呼び出し
	t2 := time.Now()
	
	
	
	return http.HandlerFunc(fn)
}