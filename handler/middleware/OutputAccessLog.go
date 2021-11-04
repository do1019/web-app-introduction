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
		ap := &model.Access{}
		start := time.Now()
		nr := SetDeviceOSInfoInContext(r)
		next.ServeHTTP(w, nr)
		end := time.Now()
		ap.Timestamp = start
		ap.Latency = int64(end.Sub(start))
		ap.Path = r.URL.Path
		osInfo, err := GetDeviceOSInfoInContext(nr)
		if err != nil {
			log.Println(err)
			return
		}
		ap.OS = osInfo
		convap, err := json.Marshal(ap)
    	if err != nil {
        	log.Println(err)
			return
    	}
    fmt.Println(string(convap))
	}
	return http.HandlerFunc(fn)
}