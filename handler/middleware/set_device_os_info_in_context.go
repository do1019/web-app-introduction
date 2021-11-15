package middleware

import (
	"context"
	"net/http"
	"github.com/mileusna/useragent"
)

type ctxKey string

const (
	OsKey ctxKey = "OS"
)

func SetDeviceOSInfoInContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		uaString := r.UserAgent()
		uaStruct := ua.Parse(uaString)
		ctx := context.WithValue(r.Context(), OsKey, uaStruct.OS)
		next.ServeHTTP(w, r.Clone(ctx))
	}
	return http.HandlerFunc(fn)
}
