package auth

import (
	"github.com/marcopeocchi/filebrowser/internal/db"
	"github.com/marcopeocchi/filebrowser/internal/domain"
)

func Container(jdb *db.JsonDB) domain.AuthHandler {
	var (
		repository = ProvideRepository(jdb)
		service    = ProvideService(repository)
		handler    = ProvideHandler(service)
	)
	return handler
}
