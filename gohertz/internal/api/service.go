package api

import "gohertz/internal/service"

type ServiceServer struct {
	exampleService service.ExampleService
	authService    service.AuthService
}

func NewService(
	exampleService service.ExampleService,
	authService service.AuthService,
) *ServiceServer {
	return &ServiceServer{
		exampleService: exampleService,
		authService:    authService,
	}
}
