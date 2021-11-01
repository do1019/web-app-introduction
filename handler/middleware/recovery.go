package middleware

import (
	"log"
	"net/http"
)

func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: ここに実装をする
		// 無名関数
		defer func() {
			if err := recover(); err != nil {
				log.Println(err);
			}
		} ()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
