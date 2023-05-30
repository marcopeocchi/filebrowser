package domain

import (
	"context"
	"net/http"
)

const (
	UserRoleAdmin = iota
	UserRoleStandard
)

type User struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type LoginForm = User

type AuthRepository interface {
	Login(ctx context.Context, username, password string) (User, error)
}

type AuthService interface {
	Login(ctx context.Context, username, password string) (User, error)
}

type AuthHandler interface {
	Login() http.HandlerFunc
	Logout() http.HandlerFunc
}
