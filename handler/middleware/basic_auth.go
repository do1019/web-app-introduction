package middleware

import (
	"fmt"
	"net/http"
	"os"
)

type AuthInfo struct {
	UserID string
	Password string
}

// 環境変数からIDとPASSを取得
func ObtainIdAndPassFromEnviron() *AuthInfo {
	ai := &AuthInfo{}
	ai.UserID = os.Getenv("BASIC_AUTH_USER_ID")
	ai.Password = os.Getenv("BASIC_AUTH_PASSWORD")
	return ai
}

// アクセス認証
func (a *AuthInfo)AccessRestriction(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if userID, pass, ok := r.BasicAuth(); userID != a.UserID || pass != a.Password || !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}