package middleware

import (
	//"fmt"
	"net/http"
	"os"
	"crypto/subtle"
	//"strings"
	//"io"
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
		if userID, pass, ok := r.BasicAuth(); !ok || subtle.ConstantTimeCompare([]byte(userID), []byte(a.UserID)) != 1 || subtle.ConstantTimeCompare([]byte(pass) , []byte(a.Password)) != 1 {
			//test
			//fmt.Println(ok)
			//fmt.Printf("userID=[%s], pass=[%s]\n", userID, pass)
			// for _, e := range os.Environ() {
			// 	pair := strings.Split(e, "=")
			// 	fmt.Println(pair[0] + "=" + pair[1])
			// }
			w.Header().Add("WWW-Authenticate", `Basic realm="my private area"`)
			//w.Write([]byte("401 認証失敗\n"))
			//
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	// 	fmt.Println("認証成功！")
	// 	io.WriteString(w, `
	// 	<!DOCTYPE html>
	// 	<html lang="ja">
	// 	<head>
	// 	  <meta charset="UTF-8">
	//    	<title>Go | net/httpパッケージ</title>
	// 	</head>
	// 	<body>
	// 　　	<h1>認証成功！</h1>
	// 　	<p>Basic認証で正しいユーザ名とパスワードが送信されました。</p>
	// 	</body>
	// 	</html>
	// 	`)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}