package middleware

import (
	//"fmt"
	"context"
	"net/http"

	"github.com/mileusna/useragent"
)

type ctx_key string

const (
	new_key ctx_key = "OS"
)

func SetDeviceOsInfoInContext(r *http.Request) *http.Request{
	uastring := r.UserAgent()
	uastruct := ua.Parse(uastring)
	ctx := context.WithValue(r.Context(), new_key, uastruct.OS)
	return r.Clone(ctx)
	// return r.Clone(ctx) //受け取り側で値が見つからない。string同士だとみつかるけど。。
}