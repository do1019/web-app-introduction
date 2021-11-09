package middleware

import (
	//"fmt"
	"context"
	"net/http"

	"github.com/do1019/web-app-introduction/handler/middleware"
	"github.com/mileusna/useragent"
)

type ctxKey string

const (
	// 先頭が小文字だと他からアクセスできない
	osKey ctxKey = "OS"
)

func SetDeviceOSInfoInContext(r *http.Request) *http.Request {
	uaString := r.UserAgent()
	uaStruct := ua.Parse(uaString)
	ctx := context.WithValue(r.Context(), osKey, uaStruct.OS)
	return r.Clone(ctx)
}

func GetDeviceOSInfoInContext(r *http.Request) (string, error) {
	ctx := r.Context()
	v := ctx.Value(osKey)
	if v == nil {
		//error構造体とメソッドをmiddlewareで定義
		return "", &middleware_error.ErrNotFound{}
	}
	osInfo, ok := v.(string)
	if !ok {
		return "", &middleware_error.ErrCannotConvType{}
	}
	return osInfo, nil
}
