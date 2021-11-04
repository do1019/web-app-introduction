package middleware

import (
	//"fmt"
	"context"
	"net/http"

	"github.com/do1019/web-app-introduction/model"
	"github.com/mileusna/useragent"
)

type ctxKey string

const (
	osKey ctxKey = "OS"
)

func SetDeviceOSInfoInContext(r *http.Request) *http.Request {
	uaString := r.UserAgent()
	uaStruct := ua.Parse(uaString)
	ctx := context.WithValue(r.Context(), osKey, uaStruct.OS)
	// 受け取り側で値が見つからない。string同士だと見つかるので型の問題。ただstringは推奨されていない。
	// ここで受け取り関数も定義してみる。
	return r.Clone(ctx)
}

func GetDeviceOSInfoInContext(r *http.Request) (string, error) {
	ctx := r.Context()
	v := ctx.Value(osKey)
	if v == nil {
		return "", &model.ErrNotFound{}
	}
	osInfo, ok := v.(string)
	if !ok {
		return "", &model.ErrCannotConvType{}
	}
	return osInfo, nil
}
