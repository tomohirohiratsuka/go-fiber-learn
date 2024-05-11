package user

import (
	"backend/internal/domain/user/handler"
	"backend/internal/domain/user/repository"
	"backend/internal/domain/user/route"
	"backend/internal/domain/user/service"
	"go.uber.org/fx"
)

//todo review fxの仕組み

// Module モジュール固有の依存関係を定義
var Module = fx.Options(
	fx.Provide(
		handler.NewUserHandler,
		fx.Annotate(service.NewUserService, fx.As(new(service.UserService))),
		fx.Annotate(repository.NewUserRepository, fx.As(new(repository.UserRepository))),
	),
	fx.Invoke(route.RegisterUserRoutes),
)
