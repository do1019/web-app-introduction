package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/do1019/web-app-introduction/handler/middleware"
)

func OutputAccessLog(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		accessInfo := &middleware.AccessInfo{}
		//Nowでtimeを取るとテストの度に結果が変わるのでやりにくい。外から入れるのが良く、ライブラリもたくさんある
		start := time.Now()
		nextRequest := SetDeviceOSInfoInContext(r)
		// SetDeviceOsInfoInContext関数を直接ServeHTTPの引数に入れたらエラーになる(?)
		next.ServeHTTP(w, nextRequest)
		end := time.Now()
		accessInfo = accessInfo.StoreTimestamp(start)
		accessInfo = accessInfo.StoreLatency(int64(end.Sub(start)))
		accessInfo = accessInfo.StorePath(r.URL.Path)
		//　storeする必要ない, os情報書き込むmiddlewareをつくる contextの意味が出てくる。
		osInfo, err := GetDeviceOSInfoInContext(nextRequest)
		// OS情報がなくてもログは出したいのでreturnしない
		if err != nil {
			log.Println(err)
		}
		accessInfo = accessInfo.StoreOS(osInfo)
		convaccessInfo, err := json.Marshal(accessInfo)
    	if err != nil {
        	log.Println(err)
			return
    	}
		fmt.Println(string(convaccessInfo))
	}
	return http.HandlerFunc(fn)
}