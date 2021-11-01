package middleware

import (
	"log"
	"net/http"
	//"fmt"
)

func Recovery(h http.Handler) http.Handler {
	//fmt.Println(1)
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: ここに実装をする
		// 無名関数, クロージャ
		//fmt.Println(3)
		defer func() {
			if err := recover(); err != nil {
				log.Println(err);
			}
		} ()
		h.ServeHTTP(w, r)
		//fmt.Println(ここまで来ない)
	}
	//fmt.Println(2)
	return http.HandlerFunc(fn)
}
