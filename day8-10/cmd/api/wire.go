//go:build wireinject
// +build wireinject

package main

import (
	"myproject/internal/db"
	"myproject/pkg/handlers"
	"myproject/pkg/models"

	"github.com/google/wire"
)

func InitializeUserHandler() (*handlers.UserHandler, error) {
	wire.Build(
		db.NewDB,
		db.NewDBConfig,
		db.NewUserRepository,
		handlers.NewUserHandler,
		wire.Bind(new(models.UserRepository), new(*db.UserRepository)),
	)
	return &handlers.UserHandler{}, nil
}
