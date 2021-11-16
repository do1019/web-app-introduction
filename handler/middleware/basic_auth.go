package middleware

import (
	"crypto/subtle"
	"net/http"
	"os"
)

type AuthInfo struct {
	UserID   string
	Password string
}

func ObtainIdAndPassFromEnviron() *AuthInfo {
	ai := &AuthInfo{}
	ai.UserID = os.Getenv("BASIC_AUTH_USER_ID")
	ai.Password = os.Getenv("BASIC_AUTH_PASSWORD")
	return ai
	// return &AuthInfo{
	// 	os.Getenv("BASIC_AUTH_USER_ID"),
	// 	os.Getenv("BASIC_AUTH_PASSWORD"),
	// }
}

func (a *AuthInfo) AccessRestriction(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userID, pass, ok := r.BasicAuth()
		if !ok ||
			subtle.ConstantTimeCompare([]byte(userID), []byte(a.UserID)) != 1 ||
			subtle.ConstantTimeCompare([]byte(pass), []byte(a.Password)) != 1 {
			// Code for testing login in a browser.
			w.Header().Add("WWW-Authenticate", `Basic realm="my private area"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
