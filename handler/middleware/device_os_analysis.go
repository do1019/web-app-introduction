package middleware

import (
	"fmt"
	//"strings"
	"net/http"

	"github.com/mileusna/useragent"
)

func Device_os_analysis(r *http.Request) {
	uastring := r.UserAgent()
	uastruct := ua.Parse(uastring)
	fmt.Println(uastruct.OS)
}