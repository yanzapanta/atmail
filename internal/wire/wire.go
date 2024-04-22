//go:build wireinject
// +build wireinject

package wire

import (
	"atmail/internal/http"
	"atmail/internal/http/handler"
	"atmail/internal/http/route"
	"atmail/internal/repository"
	"atmail/internal/service"

	"github.com/google/wire"
)

func Initialize() *http.ServerHTTP {
	wire.Build(
		route.NewUserRoute,
		handler.NewUserHandler,
		service.NewUserService,
		repository.NewUserRepository,
		http.NewServerHTTP)
	return &http.ServerHTTP{}
}
