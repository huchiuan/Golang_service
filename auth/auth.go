package auth

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("3Aju5DBXtMaFjNiANXclJ3WTcdY2pdzpi6y1USGUH0cHCLTVjZyIEQVtZbupnDGaplHgJ4/akowsthFcxWfaTRBaZZOfrpS765t+227s1W+lFPb9Yt5oeVx+yn2yOAMMIx7pCXyDM58SzdGQPEOx8+uK0gyWYhkdVgRnJIzb4iOzR6h+7mnyPfqguo3Bz9kF1u4usZRcpf4mfgB0/+G/hDe9/Y6yiqM3HqDlKQ==")

// Key cookie key
const Key = "token"

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		needAuthPaths := []string{"/list-all-users", "/get-user-detailed-information", "/search-an-user", "/delete-the-user", "/update-the-user", "/update-the-user-fullname"}
		requestPath := r.URL.Path

		var needAuth = false
		// 判斷是否是需要認證的api

		for _, path := range needAuthPaths {
			if path == requestPath {
				needAuth = true
			}
		}

		// 不需要認證，直接走下一步
		if !needAuth {
			next.ServeHTTP(w, r)
			return
		}

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "")
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "")
			return
		}
		tknStr := c.Value
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "StatusUnauthorized")
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "")
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "")
			return
		}
		next.ServeHTTP(w, r)
	})
}

type Claims struct {
	Acct string `json:"acct"`
	jwt.StandardClaims
}
