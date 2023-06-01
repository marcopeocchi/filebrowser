package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/marcopeocchi/filebrowser/internal/domain"
)

type Handler struct {
	service domain.AuthService
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		form := domain.LoginForm{}

		err := json.NewDecoder(r.Body).Decode(&form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := h.service.Login(r.Context(), form.Username, form.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userID":    user.Uid,
			"username":  user.Username,
			"role":      user.Role,
			"expiresAt": time.Now().Add(time.Hour * 24 * 30),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		cookie := http.Cookie{
			Name:     "jwt",
			HttpOnly: true,
			Secure:   false,
			Expires:  time.Now().Add(time.Hour * 24 * 30),
			Value:    tokenString,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)

		w.Write([]byte(tokenString))
	}
}

func (h *Handler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:     "jwt",
			HttpOnly: true,
			Expires:  time.Now(),
			Value:    "",
			Path:     "/",
		}
		http.SetCookie(w, &cookie)
	}
}
