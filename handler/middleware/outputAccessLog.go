package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"

	"github.com/do1019/web-app-introduction/model"
)

func OutputAccessLog(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ap := &model.AccessInfo{}
		start := time.Now()
		nr := SetDeviceOSInfoInContext(r)
		// SetDeviceOsInfoInContext関数を直接ServeHTTPの引数に入れたらエラーになる
		next.ServeHTTP(w, nr)
		end := time.Now()
		ap = ap.StoreTimestamp(start)
		ap = ap.StoreLatency(int64(end.Sub(start)))
		ap = ap.StorePath(r.URL.Path)
		osInfo, err := GetDeviceOSInfoInContext(nr)
		if err != nil {
			log.Println(err)
			return
		}
		ap = ap.StoreOS(osInfo)
		convap, err := json.Marshal(ap)
    	if err != nil {
        	log.Println(err)
			return
    	}
		fmt.Println(string(convap))
	}
	return http.HandlerFunc(fn)
}