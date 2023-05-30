package auth

import (
	"context"

	"github.com/marcopeocchi/filebrowser/internal/db"
	"github.com/marcopeocchi/filebrowser/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	jdb *db.JsonDB
}

func (r *Repository) Login(ctx context.Context, username, password string) (domain.User, error) {
	user, err := r.jdb.FindByUsername(ctx, username)
	if err != nil {
		return domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, err
	}

	return user, err
}
