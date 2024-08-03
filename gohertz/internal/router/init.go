package router

import (
	"gohertz/internal/api"
	"gohertz/internal/repository"
	"gohertz/internal/service"

	"github.com/cloudwego/hertz/pkg/app/server"
)

type routeStruct struct {
	app     *server.Hertz
	handler *api.ServiceServer
}

func NewRouter(app *server.Hertz, dao *repository.DAO) *routeStruct {
	exampleService := service.NewExampleService(dao)
	authService := service.NewAuthService(dao)
	handler := api.NewService(exampleService, authService)
	return &routeStruct{
		app:     app,
		handler: handler,
	}
}
