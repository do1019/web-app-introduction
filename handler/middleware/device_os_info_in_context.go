package middleware

import (
	"context"
	"net/http"
	"github.com/mileusna/useragent"
	"github.com/do1019/web-app-introduction/handler/middleware/error"
)

type ctxKey string

const (
	osKey ctxKey = "OS"
)

func SetDeviceOSInfoInContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		uaString := r.UserAgent()
		uaStruct := ua.Parse(uaString)
		ctx := context.WithValue(r.Context(), osKey, uaStruct.OS)
		next.ServeHTTP(w, r.Clone(ctx))
	}
	return http.HandlerFunc(fn)
}

func GetDeviceOSInfoInContext(r *http.Request) (string, error) {
	ctx := r.Context()
	v := ctx.Value(osKey)
	if v == nil {
		return "", &middleware_error.ErrNotFound{}
	}
	osInfo, ok := v.(string)
	if !ok {
		return "", &middleware_error.ErrCannotConvType{}
	}
	return osInfo, nil
}
