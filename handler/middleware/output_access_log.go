package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/do1019/web-app-introduction/model"
	"log"
	"net/http"
	"time"
)

func OutputAccessLog(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		accessInfo := &model.AccessInfo{}
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()
		accessInfo = accessInfo.StoreTimestamp(start)
		accessInfo = accessInfo.StoreLatency(int64(end.Sub(start)))
		accessInfo = accessInfo.StorePath(r.URL.Path)
		osInfo, err := GetDeviceOSInfoInContext(r)
		// Even if I can't get the OS information, I want to keep the log, so I didn't return.
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
