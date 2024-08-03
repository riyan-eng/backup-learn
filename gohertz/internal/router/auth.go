package router

import "gohertz/internal/middleware"

func (m *routeStruct) Authentication() {
	subRoute := m.app.Group("/auth")
	subRoute.POST("/register", m.handler.AuthRegister)
	subRoute.POST("/login", m.handler.AuthLogin)
	subRoute.POST("/refresh", m.handler.AuthRefresh)
	subRoute.Use(middleware.Jwt())
	subRoute.GET("/me", m.handler.AuthMe)
}
