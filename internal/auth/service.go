package auth

import (
	"context"

	"github.com/marcopeocchi/filebrowser/internal/domain"
)

type Service struct {
	repository domain.AuthRepository
}

func (s *Service) Login(ctx context.Context, username, password string) (domain.User, error) {
	return s.repository.Login(ctx, username, password)
}
