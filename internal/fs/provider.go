package fs

import (
	"sync"

	"github.com/marcopeocchi/filebrowser/internal/domain"
)

var (
	repository     *Repository
	service        *Service
	handler        *Handler
	repositoryOnce sync.Once
	serviceOnce    sync.Once
	handlerOnce    sync.Once
)

func ProvideRepository(root string) *Repository {
	repositoryOnce.Do(func() {
		repository = &Repository{
			root: root,
		}
	})
	return repository
}

func ProvideService(r domain.FSRepository) *Service {
	serviceOnce.Do(func() {
		service = &Service{
			repository: r,
		}
	})
	return service
}

func ProvideHandler(s domain.FSService) *Handler {
	handlerOnce.Do(func() {
		handler = &Handler{
			service: s,
		}
	})
	return handler
}
