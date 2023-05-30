package auth

import (
	"sync"

	"github.com/marcopeocchi/filebrowser/internal/db"
	"github.com/marcopeocchi/filebrowser/internal/domain"
)

var (
	repository *Repository
	service    *Service
	handler    *Handler

	repositoryOnce sync.Once
	serviceOnce    sync.Once
	handlerOnce    sync.Once
)

func ProvideRepository(jdb *db.JsonDB) *Repository {
	repositoryOnce.Do(func() {
		repository = &Repository{
			jdb: jdb,
		}
	})
	return repository
}

func ProvideService(repo domain.AuthRepository) *Service {
	serviceOnce.Do(func() {
		service = &Service{
			repository: repo,
		}
	})
	return service
}

func ProvideHandler(svc domain.AuthService) *Handler {
	handlerOnce.Do(func() {
		handler = &Handler{
			service: svc,
		}
	})
	return handler
}
