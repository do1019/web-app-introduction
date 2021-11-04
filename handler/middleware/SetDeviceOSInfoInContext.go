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

func SetDeviceOSInfoInContext(r *http.Request) *http.Request {
	uastring := r.UserAgent()
	uastruct := ua.Parse(uastring)
	ctx := context.WithValue(r.Context(), new_key, uastruct.OS)
	// 受け取り側で値が見つからない。string同士だと見つかるので型の問題。ただstringは推奨されていない。
	// ここで受け取り関数も定義してみる。
	return r.Clone(ctx)
}

func GetDeviceOSInfoInContext(r *http.Request) (string, error) {
	ctx := r.Context()
	if v := ctx.Value(new_key); v == nil {
		return "", "NotFound"
	}
}