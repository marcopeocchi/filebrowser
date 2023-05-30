package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")

		if err != nil {
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		if cookie == nil {
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		token, _ := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("JWTSECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expiresAt, err := time.Parse(time.RFC3339, claims["expiresAt"].(string))

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if time.Now().After(expiresAt) {
				//http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
				http.Error(w, "token expired", http.StatusBadRequest)
				return
			}
		} else {
			//http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
