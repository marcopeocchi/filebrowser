package fs

import (
	"context"

	"github.com/marcopeocchi/filebrowser/internal/domain"
)

type Service struct {
	repository domain.FSRepository
}

func (s *Service) WalkDir(ctx context.Context, subDir string) (domain.Response, error) {
	return s.repository.WalkDir(ctx, subDir)
}

func (s *Service) GetBasePathLength(ctx context.Context) (int, error) {
	return s.repository.GetBasePathLength(ctx)
}
