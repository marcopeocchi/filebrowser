package fs

import "github.com/marcopeocchi/filebrowser/internal/domain"

func Container(root string) domain.FSHandler {
	var (
		repository = ProvideRepository(root)
		service    = ProvideService(repository)
		handler    = ProvideHandler(service)
	)
	return handler
}
